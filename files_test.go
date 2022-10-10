package main

import (
	"fmt"
	"testing"
)

func TestProcessFiles(t *testing.T) {

	files := make([]string, 2)
	files[0] = "file1"
	files[1] = "file2"

	var noFiles []string

	var tests = []struct {
		inputPercentage int
		inputFiles      []string
		expectedResult  bool
	}{
		{0, files, true},
		{10, noFiles, true},
		{-10, files, true},
		{10, files, true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("Percentage: %d Files: %s", currentTest.inputPercentage, currentTest.inputFiles)
		t.Run(testname, func(t *testing.T) {
			result := ProcessFiles(currentTest.inputPercentage, currentTest.inputFiles)
			if result != currentTest.expectedResult {
				t.Errorf("ProcessFiles() : FAIL : input %d %s got %t, want %t", currentTest.inputPercentage, currentTest.inputFiles, result, currentTest.expectedResult)
			} else {
				t.Logf("ProcessFiles() : PASS : input %d %s got %t, want %t", currentTest.inputPercentage, currentTest.inputFiles, result, currentTest.expectedResult)
			}
		})
	}
}

func TestFileExists(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"fakeFile", false},
		{"psrm", true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := FileExists(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("FileExists() : FAIL : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("FileExists() : PASS : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestFileUser(t *testing.T) {

	var tests = []struct {
		input                   string
		expectedResultUser      string
		expectedResultUserFound bool
	}{
		{"", "", false},
		{"fakeFile", "", false},
		{"psrm", "pi", true},
		{"test data/rootFile", "root", true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			actualResultUser, actualResultUserFound := FileUser(currentTest.input)
			if actualResultUser != currentTest.expectedResultUser || actualResultUserFound != currentTest.expectedResultUserFound {
				t.Errorf("FileUser() : FAIL : input %s got %s %t, want %s %t", currentTest.input, actualResultUser, actualResultUserFound, currentTest.expectedResultUser, currentTest.expectedResultUserFound)
			} else {
				t.Logf("FileUser() : PASS : input %s got %s %t, want %s %t", currentTest.input, actualResultUser, actualResultUserFound, currentTest.expectedResultUser, currentTest.expectedResultUserFound)
			}
		})
	}
}

func TestLinuxFileUser(t *testing.T) {

	var tests = []struct {
		input                   string
		expectedResultUser      string
		expectedResultUserFound bool
	}{
		{"", "", false},
		{"fakeFile", "", false},
		{"psrm", "pi", true},
		{"test data/rootFile", "root", true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			actualResultUser, actualResultUserFound := LinuxFileUser(currentTest.input)
			if actualResultUser != currentTest.expectedResultUser || actualResultUserFound != currentTest.expectedResultUserFound {
				t.Errorf("LinuxFileUser() : FAIL : input %s got %s %t, want %s %t", currentTest.input, actualResultUser, actualResultUserFound, currentTest.expectedResultUser, currentTest.expectedResultUserFound)
			} else {
				t.Logf("LinuxFileUser() : PASS : input %s got %s %t, want %s %t", currentTest.input, actualResultUser, actualResultUserFound, currentTest.expectedResultUser, currentTest.expectedResultUserFound)
			}
		})
	}
}

func TestFileWriteable(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"fakeFile", false},
		{"psrm", true},
		{"test data/rootFile", true},
		{"test data/readonly", false},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := FileWriteable(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("FileWriteable() : FAIL : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("FileWriteable() : PASS : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestFileWriteableLinux(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"fakeFile", false},
		{"psrm", true},
		{"test data/rootFile", true},
		{"test data/readonly", false},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := FileWriteableLinux(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("FileWriteableLinux() : FAIL : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("FileWriteableLinux() : PASS : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestGetFileSize(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult int64
	}{
		{"", -1},
		{"fakeFile", -1},
		{"test data/mediumFile", 18},
		{"test data/smallFile", 6},
		{"test data/emptyFile", -1},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := GetFileSize(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("GetFileSize() : FAIL : input %s got %d, want %d", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("GetFileSize() : PASS : input %s got %d, want %d", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestOverwriteFile(t *testing.T) {

	var tests = []struct {
		inputPercentage int
		inputFile       string
		expectedResult  bool
	}{
		{0, "fakeFile1", true},
		{10, "", true},
		{-10, "fakeFile1", true},
		{10, "fakeFile2", true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("Percentage: %d File: %s", currentTest.inputPercentage, currentTest.inputFile)
		t.Run(testname, func(t *testing.T) {
			result := OverwriteFile(currentTest.inputPercentage, currentTest.inputFile)
			if result != currentTest.expectedResult {
				t.Errorf("OverwriteFile() : FAIL : input %d %s got %t, want %t", currentTest.inputPercentage, currentTest.inputFile, result, currentTest.expectedResult)
			} else {
				t.Logf("OverwriteFile() : PASS : input %d %s got %t, want %t", currentTest.inputPercentage, currentTest.inputFile, result, currentTest.expectedResult)
			}
		})
	}
}

func TestWriteToFile(t *testing.T) {

	var tests = []struct {
		inputFile      string
		inputSections  []section
		expectedResult bool
	}{
		{"", make([]section, 2), true},
		{"", nil, true},
		{"fakeFile1", make([]section, 2), true},
		{"test data/rootFile", nil, true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("File: %s", currentTest.inputFile)
		t.Run(testname, func(t *testing.T) {
			result := WriteToFile(currentTest.inputFile, currentTest.inputSections)
			if result != currentTest.expectedResult {
				t.Errorf("WriteToFile() : FAIL : input %s got %t, want %t", currentTest.inputFile, result, currentTest.expectedResult)
			} else {
				t.Logf("WriteToFile() : PASS : input %s got %t, want %t", currentTest.inputFile, result, currentTest.expectedResult)
			}
		})
	}
}

func TestDeleteFile(t *testing.T) {

	var tests = []struct {
		inputFile      string
		expectedResult bool
	}{
		{"", true},
		{"fakeFile1", true},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("File: %s", currentTest.inputFile)
		t.Run(testname, func(t *testing.T) {
			result := DeleteFile(currentTest.inputFile)
			if result != currentTest.expectedResult {
				t.Errorf("WriteToFile() : FAIL : input %s got %t, want %t", currentTest.inputFile, result, currentTest.expectedResult)
			} else {
				t.Logf("WriteToFile() : PASS : input %s got %t, want %t", currentTest.inputFile, result, currentTest.expectedResult)
			}
		})
	}
}
