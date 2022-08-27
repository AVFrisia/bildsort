package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// The folder to which we will write images
var OutPath = "out"

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s [srcDir] [dstDir]\n", os.Args[0])
		os.Exit(1)
	}

	src := os.Args[1]
	dst := os.Args[2]

	// Use absolute paths to prevent any sort of confusion
	src, err := filepath.Abs(src)
	if err != nil {
		log.Fatal(err)
	}

	dst, err = filepath.Abs(dst)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sorting from %s to %s\n", src, dst)

	OutPath = dst

	err = filepath.WalkDir(src, processImage)
	
	if err != nil {
		log.Print(err)
	}
}
