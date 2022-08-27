package main

import (
	"testing"
	"time"
)

func testMonths(t *testing.T, year, startMonth, endMonth int, expect string) {
	for m := startMonth; m < endMonth+1; m++ {
		for d := 1; d < 30; d++ {
			dat := time.Date(year, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
			res := GetSemester(dat)
			if res != expect {
				t.Fatalf("%s is %s, should be %s", dat, res, expect)
			}
		}
	}
}

func TestSummer(t *testing.T) {
	exp := "SoSe 2016"
	testMonths(t, 2016, 4, 9, exp)
}

func TestWinter(t *testing.T) {
	exp := "WiSe 2016 - 2017"
	testMonths(t, 2016, 10, 12, exp)
	testMonths(t, 2017, 1, 3, exp)
}
