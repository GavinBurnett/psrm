package main

import (
	"fmt"
	"math"
	"unicode"
)

// Returns true if given string is a number
func IsStringNumber(_theString string) bool {

	isString := true

	if len(_theString) > 0 {

		for _, currentChar := range _theString {
			if !unicode.IsDigit(currentChar) {
				isString = false
				break
			}
		}

	} else {
		isString = false
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_theString:", _theString)
	}

	return isString
}

// Returns true if given string is a help argument
func IsStringHelpArgument(_theString string) bool {

	isHelpArgument := false

	if len(_theString) > 0 {

		switch _theString {
		case "?":
			isHelpArgument = true
		case "/?":
			isHelpArgument = true
		case "-?":
			isHelpArgument = true
		case "--?":
			isHelpArgument = true
		case "h":
			isHelpArgument = true
		case "/h":
			isHelpArgument = true
		case "-h":
			isHelpArgument = true
		case "--h":
			isHelpArgument = true
		case "help":
			isHelpArgument = true
		case "/help":
			isHelpArgument = true
		case "-help":
			isHelpArgument = true
		case "--help":
			isHelpArgument = true
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_theString:", _theString)
	}

	return isHelpArgument
}

// Display section type name
func DisplaySectionType(_sectionType SECTION_TYPE) string {

	sectionType := ""

	if _sectionType >= 1 && _sectionType <= math.MaxInt16 {

		switch _sectionType {
		case RANDOM_SPACE:
			sectionType = UI_RandomSpace
		case OVERWRITE_SECTION:
			sectionType = UI_OverwriteSection
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_sectionType:", _sectionType)
	}

	return sectionType
}
