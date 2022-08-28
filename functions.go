package main

import (
	"errors"
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	"golang.org/x/exp/slices"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

const DateTimeId = 0x9003
const ExifDateLayout = "2006:01:02 15:04:05"

// GetSemester parses a time and returns the appropriate semester.
//
// According to the Leibniz University of Hannover, these are:
// Wintersemester: 1st October -- 31. March
// Summersemester: 1st April -- 30. September
//
// Names are shortened to semester and year.
// For Example: WiSe 2016 - 2017, SoSe 2019 or WiSe 2022 - 2023
func GetSemester(date time.Time) string {
	year := date.Year()

	// Define limits of the  that year's summer semester
	sumStart := time.Date(year, 4, 1, 0, 0, 0, 0, time.UTC)
	sumEnd := time.Date(year, 9, 30, 0, 0, 0, 0, time.UTC)

	// Check if in limits for summer
	if date.After(sumEnd) {
		return fmt.Sprintf("WiSe %d - %d", year, year+1)
	} else if date.Before(sumStart) {
		return fmt.Sprintf("WiSe %d - %d", year-1, year)
	} else {
		return fmt.Sprintf("SoSe %d", year)
	}
}

// Move images to a matching folder
func move(path string, d fs.DirEntry, date time.Time) error {
	sem := GetSemester(date)
	targetDir := filepath.Join(OutPath, sem)

	err := os.MkdirAll(targetDir, 0750)
	if err != nil {
		return err
	}

	targetPath := filepath.Join(targetDir, d.Name())

	// COPY file
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(targetPath, data, 0750)
	if err != nil {
		return err
	}

	log.Printf("%s -> %s\n", path, targetPath)
	return nil
}

// allowed simply checks to see if a given file should be processed
// (i.e. it is one of the recognized image files)
func allowed(d fs.DirEntry) bool {
	// Entry must be file
	if d.IsDir() {
		return false
	}

	// TODO: Perform some kind of pattern matching
	// For now, files without EXIF (non-Images) fail silently
	return true
}

// ExifDate finds and extracts a TimeStamp from an array of EXIF tags.
func ExifDate(exifTags []exif.ExifTag) (time.Time, error) {
	ti := slices.IndexFunc(exifTags, func(t exif.ExifTag) bool { return t.TagId == DateTimeId })

	if ti < 0 {
		return time.Time{}, errors.New("File has no timestamp!")
	}

	dateString := exifTags[ti].Formatted

	timeStamp, err := time.Parse(ExifDateLayout, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return timeStamp, nil
}

func LastModified(path string) (time.Time, error) {
	file, err := os.Stat(path)

	if err != nil {
		return time.Time{}, err
	}

	return file.ModTime(), nil
}

// ExtractDate attempts to find a date for an image file using its EXIF metadata.
//
// If no EXIF tags are available, it uses the file's modification date.
func ExtractDate(path string) (time.Time, error) {
	data, err := exif.SearchFileAndExtractExif(path)
	if err != nil {
		log.Print(err, "using last modified")
		return LastModified(path)
	}

	tags, _, err := exif.GetFlatExifData(data, nil)
	if err != nil {
		log.Print(err)
		return LastModified(path)
	}

	date, err := ExifDate(tags)
	if err != nil {
		log.Print(err)
		return LastModified(path)
	}

	return date, nil
}

func processImage(path string, d fs.DirEntry, err error) error {
	if !allowed(d) {
		return nil
	}

	date, err := ExtractDate(path)
	if err != nil {
		log.Print(err)
		return nil
	}

	err = move(path, d, date)
	if err != nil {
		log.Print(err)
		return nil
	}

	return nil
}
