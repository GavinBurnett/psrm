package main

import (
	"fmt"
	"testing"
)

func TestFileUserAccess(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"fakeFile", false},
		{"psrm", true},
		{"test data/rootFile", false},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := FileUserAccess(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("FileUserAccess() : FAIL : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("FileUserAccess() : PASS : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestFileUserAccessLinux(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"fakeFile", false},
		{"psrm", true},
		{"test data/rootFile", false},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := FileUserAccessLinux(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("FileUserAccessLinux() : FAIL : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("FileUserAccessLinux() : PASS : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}
