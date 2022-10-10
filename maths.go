package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Defines sections of file to process
type section struct {
	sectionType SECTION_TYPE
	start       int64
	end         int64
	size        int64
}

// Returns true if given value is in range of 1 to 100, false if outside of range
func IsPercentageValid(_percentage int) bool {

	percentageValid := false

	if _percentage > 0 && _percentage <= math.MaxInt32 {

		if _percentage >= 1 && _percentage <= 100 {
			percentageValid = true
		} else {
			percentageValid = false
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_percentage:", _percentage)
	}

	return percentageValid
}

// Gets the size of data (in bytes) to overwrite for the given file
func GetFilePercentage(_fileSize int64, _percentage int) int64 {

	var filePercentageSize int64
	filePercentageSize = -1

	if _fileSize > 0 && _fileSize <= math.MaxInt64 && _percentage > 0 && _percentage <= math.MaxInt32 {

		if _fileSize <= 100 {
			filePercentageSize = _fileSize
		} else {
			filePercentageSize = _fileSize * int64(_percentage) / 100
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_fileSize:", _fileSize, "_percentage:", _percentage)
	}

	return filePercentageSize
}

// Gets a single file section to overwrite entire file
func GetSection(_fileSize int64) []section {

	var sections []section
	var currentSection section

	if _fileSize >= 1 && _fileSize <= math.MaxInt64 {

		currentSection.sectionType = OVERWRITE_SECTION
		currentSection.start = 0
		currentSection.end = _fileSize
		currentSection.size = _fileSize

		sections = append(sections, currentSection)

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_fileSize:", _fileSize)
	}

	return sections
}

// Gets the file locations for the sections to overwrite
func GetSections(_fileSize int64, _filePercentageToOverwriteSize int64) []section {

	var sections []section
	sections = nil

	var sectionsFit = false

	if _fileSize >= 1 && _fileSize <= math.MaxInt64 && _filePercentageToOverwriteSize >= 1 && _filePercentageToOverwriteSize <= math.MaxInt64 {

		if _filePercentageToOverwriteSize > 100 {

			var sectionSize int64
			sectionSize = -1

			numberOfSections := -1

			sectionSize, numberOfSections = GetSectionSize(_filePercentageToOverwriteSize, SECTIONS)

			if sectionSize != -1 && numberOfSections != -1 {

				var offset int64
				offset = 0

				overwriteSection := false

				for checkCounter := 0; checkCounter != 5 && sectionsFit == false; checkCounter++ {

					for counter := 0; counter != numberOfSections; counter++ {

						var currentSection section

						if overwriteSection == true {
							currentSection.sectionType = OVERWRITE_SECTION
							currentSection.start = offset
							currentSection.end = offset + sectionSize
							currentSection.size = sectionSize
						} else {
							currentSection.sectionType = RANDOM_SPACE
							currentSection.start = offset

							currentSection.end = offset + GetRandomNumber(int64(numberOfSections)*10, int64(numberOfSections)*500)
						}

						offset = currentSection.end

						// if sections are greater than file size - stop
						if offset > _fileSize {
							break
						}

						if overwriteSection == true {
							overwriteSection = false
						} else {
							overwriteSection = true
						}

						sections = append(sections, currentSection)
					} // end sections loop

					sectionsFit = CheckSections(_fileSize, sections)

				} // end check sections loop
			}
		}

		if sectionsFit == false {
			sections = nil
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_fileSize:", _fileSize, "_filePercentageToOverwriteSize:", _filePercentageToOverwriteSize)
	}

	return sections
}

// Check that file offsets are not larger than the file size
func CheckSections(_fileSize int64, _sections []section) bool {

	pass := true

	if _fileSize <= math.MaxInt64 && _sections != nil && len(_sections) > 0 {

		if _fileSize < 1 {
			pass = false
		} else {
			if _fileSize >= 1 && _fileSize <= math.MaxInt64 && len(_sections) > 0 {
				for counter := 0; counter != len(_sections); counter++ {
					if _sections[counter].start >= _fileSize || _sections[counter].end >= _fileSize {
						pass = false
						break
					}
				} // end loop
			}
		}

	} else {
		pass = false
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_fileSize:", _fileSize, "_sections:", _sections)
	}

	return pass
}

// Gets the size of a single section in bytes, and number of sections
func GetSectionSize(_filePercentageToOverwrite int64, _numberOfSections int64) (int64, int) {

	var sectionSize int64
	sectionSize = -1

	var numberOfSections int
	numberOfSections = -1

	if _filePercentageToOverwrite >= 1 && _filePercentageToOverwrite <= math.MaxInt64 && _numberOfSections >= 1 && _numberOfSections <= math.MaxInt64 {
		for currentSection := _numberOfSections; currentSection != 0; currentSection-- {
			sectionSize = _filePercentageToOverwrite / _numberOfSections
			if sectionSize <= 0 {
				if currentSection == 1 && sectionSize == 0 {
					sectionSize = 1
					numberOfSections = int(currentSection)
				}
			} else {
				numberOfSections = int(currentSection)
				break
			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_filePercentageToOverwrite:", _filePercentageToOverwrite, "_numberOfSections:", _numberOfSections)
	}

	return sectionSize, numberOfSections
}

// Gets a random number between given low and high values
func GetRandomNumber(_low int64, _high int64) int64 {

	var randomNumber int64
	randomNumber = -1

	if _low >= 1 && _low <= math.MaxInt64 && _high >= 1 && _high <= math.MaxInt64 {
		if _low < _high {
			rand.Seed(time.Now().UnixNano())
			randomNumber = rand.Int63n(_high-_low+1) + _low
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_low:", _low, "_high:", _high)
	}

	return randomNumber
}

// Gets a byte array of given size filled with random data
func GetRandomData(_size int64) []byte {

	var randomData []byte
	randomData = nil

	if _size > 0 && _size <= math.MaxInt64 {

		randomData = make([]byte, _size)

		if int64(len(randomData)) != _size {
			fmt.Println(UI_RandomDataError)
		} else {
			_, err := rand.Read(randomData)

			if err != nil {
				fmt.Println(UI_RandomDataError, _size, err.Error())
			}

			if int64(len(randomData)) != _size {
				fmt.Println(UI_RandomDataError)
			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_size:", _size)
	}

	return randomData
}
