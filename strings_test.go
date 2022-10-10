package main

import (
	"fmt"
	"testing"
)

func TestIsStringNumber(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"-10", false},
		{"0", true},
		{"10", true},
		{"100", true},
		{"abc", false},
		{"1a", false},
		{"a1", false},
		{"1a1", false},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := IsStringNumber(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("IsStringNumber() : FAIL : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("IsStringNumber() : PASS : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestIsStringHelpArgument(t *testing.T) {

	var tests = []struct {
		input          string
		expectedResult bool
	}{
		{"", false},
		{"?", true},
		{"/?", true},
		{"-?", true},
		{"--?", true},
		{"-h", true},
		{"--h", true},
		{"help", true},
		{"/help", true},
		{"-help", true},
		{"--help", true},
		{"a", false},
		{"12", false},
		{"+", false},
		{"-a", false},
		{"--a", false},
		{"/*", false},
		{"-9", false},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.input)
		t.Run(testname, func(t *testing.T) {
			result := IsStringHelpArgument(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("IsStringHelpArgument() : FAIL : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			} else {
				t.Logf("IsStringHelpArgument() : PASS : input %s got %t, want %t", currentTest.input, result, currentTest.expectedResult)
			}
		})
	}
}

func TestDisplaySectionType(t *testing.T) {

	var tests = []struct {
		input          SECTION_TYPE
		expectedResult string
	}{
		{OVERWRITE_SECTION, UI_OverwriteSection},
		{RANDOM_SPACE, UI_RandomSpace},
	}

	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", string(currentTest.input))
		t.Run(testname, func(t *testing.T) {
			result := DisplaySectionType(currentTest.input)
			if result != currentTest.expectedResult {
				t.Errorf("DisplaySectionType() : FAIL : input %s got %s, want %s", string(currentTest.input), result, currentTest.expectedResult)
			} else {
				t.Logf("DisplaySectionType() : PASS : input %s got %s, want %s", string(currentTest.input), result, currentTest.expectedResult)
			}
		})
	}
}
