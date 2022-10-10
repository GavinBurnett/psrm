package main

import (
	"fmt"
	"testing"
)

func TestIsPercentageValid(t *testing.T) {

	var tests = []struct {
		input          int
		expectedResult bool
	}{
		{-20, false},
		{0, false},
		{1, true},
		{20, true},
		{99, true},
		{100, true},
		{101, false},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%d", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := IsPercentageValid(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("IsPercentageValid() : FAIL : input %d got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("IsPercentageValid() : PASS : input %d got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestGetFilePercentage(t *testing.T) {

	var tests = []struct {
		input_fileSize   int64
		input_percentage int
		expectedResult   int64
	}{
		{0, 10, -1},
		{1, 0, -1},
		{1, 10, 1},
		{50, 20, 50},
		{100, 100, 100},
		{1000, 1, 10},
		{1000, 10, 100},
		{1000, 75, 750},
		{1000, 100, 1000},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("FileSize: %d Percentage: %d", currentTest.input_fileSize, currentTest.input_percentage)
		t.Run(testname, func(t *testing.T) {
			result := GetFilePercentage(currentTest.input_fileSize, currentTest.input_percentage)
			if result != currentTest.expectedResult {
				t.Errorf("GetFilePercentage() : FAIL : input: FileSize: %d Percentage: %d - got %d, want %d", currentTest.input_fileSize, currentTest.input_percentage, result, currentTest.expectedResult)
			} else {
				t.Logf("GetFilePercentage() : PASS : input: FileSize: %d Percentage: %d - got %d, want %d", currentTest.input_fileSize, currentTest.input_percentage, result, currentTest.expectedResult)
			}
		})
	}
}

func TestGetSection(t *testing.T) {

	var currentSection section
	currentSection.sectionType = OVERWRITE_SECTION

	var tests = []struct {
		input_fileSize int64
		expectedResult section
	}{
		{1, currentSection},
		{50, currentSection},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("FileSize: %d", currentTest.input_fileSize)
		t.Run(testname, func(t *testing.T) {
			result := GetSection(currentTest.input_fileSize)
			if result[0].sectionType != currentTest.expectedResult.sectionType {
				t.Errorf("GetSection() : FAIL : input: FileSize: %d", currentTest.input_fileSize)
			} else {
				t.Logf("GetSection() : PASS : input: FileSize: %d", currentTest.input_fileSize)
			}
		})
	}
}

func TestGetSections(t *testing.T) {

	var tests = []struct {
		input_fileSize                      int64
		input_filePercentageToOverwriteSize int64
		expectedResult                      int
	}{
		{0, 200, 0},
		{500000, 0, 0},
		{1, 1, 0},
		{50000, 5, 0},
		{50000, 200, 10},
		{500000, 200, 10},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("FileSize: %d", currentTest.input_fileSize)
		t.Run(testname, func(t *testing.T) {
			var sections []section
			sections = GetSections(currentTest.input_fileSize, currentTest.input_filePercentageToOverwriteSize)
			if len(sections) != currentTest.expectedResult {
				t.Errorf("GetSections() : FAIL : input: FileSize: %d Percentage: %d - got %d, want %d", currentTest.input_fileSize, currentTest.input_filePercentageToOverwriteSize, len(sections), currentTest.expectedResult)
			} else {
				t.Logf("GetSections() : PASS : input: FileSize: %d Percentage: %d - got %d, want %d", currentTest.input_fileSize, currentTest.input_filePercentageToOverwriteSize, len(sections), currentTest.expectedResult)
			}
		})
	}
}

func TestCheckSections(t *testing.T) {

	var sections []section
	for counter := 0; counter != 3; counter++ {
		var currentSection section
		currentSection.start = 10
		currentSection.end = 20
		sections = append(sections, currentSection)
	}

	var tests = []struct {
		input_filesize int64
		expectedResult bool
	}{
		{-1, false},
		{0, false},
		{1, false},
		{10, false},
		{50, true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%d", currentTest.input_filesize)
		t.Run(testname, func(t *testing.T) {
			sectionsFit := CheckSections(currentTest.input_filesize, sections)
			if sectionsFit == currentTest.expectedResult {
				t.Logf("CheckSections() : PASS : input %d got %t, want %t", currentTest.input_filesize, sectionsFit, currentTest.expectedResult)
			} else {
				t.Errorf("CheckSections() : FAIL : input %d got %t, want %t", currentTest.input_filesize, sectionsFit, currentTest.expectedResult)
			}
		})
	}
}

func TestGetSectionSize(t *testing.T) {

	var tests = []struct {
		input                           int64
		expectedResult_sectionSize      int64
		expectedResult_numberOfSections int
	}{
		{-10, -1, -1},
		{0, -1, -1},
		{1, 1, 1},
		{5, 1, 1},
		{10, 1, 10},
		{15, 1, 10},
		{20, 2, 10},
		{30, 3, 10},
		{40, 4, 10},
		{50, 5, 10},
		{100, 10, 10},
		{1000, 100, 10},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%d", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result, result2 := GetSectionSize(currentTest.input, SECTIONS)
			if result != currentTest.expectedResult_sectionSize || result2 != currentTest.expectedResult_numberOfSections {
				t.Errorf("GetSectionSize() : FAIL : input %d got %d - %d, want %d - %d", currentTest.input, result, result2, currentTest.expectedResult_sectionSize, currentTest.expectedResult_numberOfSections)
			} else {
				t.Logf("GetSectionSize() : PASS : input %d got %d - %d, want %d - %d", currentTest.input, result, result2, currentTest.expectedResult_sectionSize, currentTest.expectedResult_numberOfSections)
			}
		})
	}
}

func TestGetRandomNumber(t *testing.T) {

	var tests = []struct {
		input_low      int64
		input_high     int64
		expectedResult int64
	}{
		{0, 0, -1},
		{0, -1, -1},
		{-1, 0, -1},
		{-1, -1, -1},
		{1, 1, -1},
		{10, 10, -1},
		{2, 1, -1},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("Low: %d High: %d", currentTest.input_low, currentTest.input_high)
		t.Run(testname, func(t *testing.T) {
			result := GetRandomNumber(currentTest.input_low, currentTest.input_high)
			if result != currentTest.expectedResult {
				t.Errorf("GetSectionSize() : FAIL : input %d - %d got %d, want %d", currentTest.input_low, currentTest.input_high, result, currentTest.expectedResult)
			} else {
				t.Logf("GetSectionSize() : PASS : input %d - %d got %d, want %d", currentTest.input_low, currentTest.input_high, result, currentTest.expectedResult)
			}
		})
	}
}

func TestGetRandomData(t *testing.T) {

	var tests = []struct {
		input          int64
		expectedResult int64
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{10, 10},
		{100, 100},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("Input: %d", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := GetRandomData(currentTest.input)
			if int64(len(result)) != currentTest.expectedResult {
				t.Errorf("TestRandomData() : FAIL : input %d got data: %d - size: %d, want size: %d", currentTest.input, result, len(result), currentTest.expectedResult)
			} else {
				t.Logf("TestRandomData() : PASS : input %d got data: %d - size: %d, want size: %d", currentTest.input, result, len(result), currentTest.expectedResult)
			}
		})
	}
}
