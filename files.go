package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// Overwrite passed in percentge of all passed in file/files with random data, then delete the file/files
func ProcessFiles(_percentage int, _files []string) bool {

	if DEBUG == true {
		fmt.Println("All arguments: ", _percentage, _files)
	}

	overwriteError := false
	deleteError := false

	processFileError := false

	if _percentage > 0 && _percentage <= math.MaxInt32 && _files != nil && len(_files) > 0 {

		for _, file := range _files {
			overwriteError = false
			deleteError = false
			if FileExists(file) {
				if FileWriteable(file) {
					if FileUserAccess(file) {
						fmt.Printf(file)
						overwriteError = OverwriteFile(_percentage, file)
						if overwriteError == false {
							deleteError = DeleteFile(file)
							if deleteError == false {
								fmt.Println(UI_Deleted)
							} else {
								processFileError = true
								fmt.Println(UI_DeleteError, file)
							}
						} else {
							processFileError = true
							fmt.Println(UI_OverwriteError, file)
						}
					} else {
						processFileError = true
						fmt.Println(UI_FileNoUserPerm, file)
					}
				} else {
					processFileError = true
					fmt.Println(UI_FileNoWritePerm, file)
				}
			} else {
				processFileError = true
				fmt.Println(UI_FileNotFound, file)
			}
		}
	} else {
		processFileError = true
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_percentage:", _percentage, "_files:", _files)
	}

	if overwriteError == true || deleteError == true {
		processFileError = true
	}

	return processFileError
}

// Check the given file exists
func FileExists(_file string) bool {

	fileExists := false

	if len(_file) > 0 {
		_, err := os.Stat(_file)
		if os.IsNotExist(err) || err != nil {
			fileExists = false
		} else {
			fileExists = true
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return fileExists
}

// Returns user of passed in file
func FileUser(_file string) (string, bool) {

	fileUser := ""
	fileUserFound := false

	if len(_file) > 0 {

		switch PLATFORM {
		case PI_ARM_PLATFORM:
			fileUser, fileUserFound = LinuxFileUser(_file)
		case LINUX_PLATFORM:
			fileUser, fileUserFound = LinuxFileUser(_file)
		case WINDOWS_PLATFORM:
			fileUser = "root"
			fileUserFound = true
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return fileUser, fileUserFound
}

// Returns user of passed in file - linux platform
func LinuxFileUser(_file string) (string, bool) {

	fileUser := ""
	fileUserFound := false

	if len(_file) > 0 {

		fileInfo, err := os.Stat(_file)

		if err != nil {
			fmt.Println(UI_FileUserError, err.Error())
		} else {
			stat := fileInfo.Sys().(*syscall.Stat_t)
			if stat != nil {
				uid := int(stat.Uid)
				userStr := strconv.FormatUint(uint64(uid), 10)

				if userData, err := user.LookupId(userStr); err != nil {
					fmt.Println(UI_FileUserError, _file, err.Error())
				} else {
					fileUserFound = true
					fileUser = userData.Username
				}
			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return fileUser, fileUserFound
}

// Check the given file can be written to
func FileWriteable(_file string) bool {

	writePermission := false

	if len(_file) > 0 {

		switch PLATFORM {
		case PI_ARM_PLATFORM:
			writePermission = FileWriteableLinux(_file)
		case LINUX_PLATFORM:
			writePermission = FileWriteableLinux(_file)
		case WINDOWS_PLATFORM:
			writePermission = true
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return writePermission
}

// Check the given file can be written to - linux
func FileWriteableLinux(_file string) bool {

	writePermission := false

	if len(_file) > 0 {

		fileInfo, err := os.Stat(_file)

		if err == nil {
			fileMode := fileInfo.Mode()

			if string(fileMode.String()[2]) == "w" {
				writePermission = true
			} else {
				writePermission = false
			}
		} else {
			fmt.Println(UI_FileStatError, _file, err.Error())
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return writePermission
}

// Gets the size of the given file in bytes
func GetFileSize(_file string) int64 {

	var fileSize int64
	fileSize = -1

	if len(_file) > 0 {

		fileInfo, err := os.Stat(_file)
		if err == nil {
			fileSize = fileInfo.Size()

			if fileSize == 0 {
				fmt.Println(UI_EmptyFile)
				fileSize = -1
			}

			if fileSize < 0 {
				fmt.Println(UI_InvalidFileSize)
				fileSize = -1
			}

			if fileSize > math.MaxInt64 {
				fmt.Println(UI_FileTooBig)
				fileSize = -1
			}

		} else {
			fmt.Println(UI_NoFileSize, _file, err.Error())
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return fileSize
}

// Overwrite the given percentage of the given file with random data
func OverwriteFile(_percentage int, _file string) bool {

	overwriteError := false

	if _percentage > 0 && _percentage <= math.MaxInt32 && len(_file) > 0 {

		fileSize := GetFileSize(_file)
		if fileSize != -1 {
			filePercentageToOverwriteSize := GetFilePercentage(fileSize, _percentage)
			if filePercentageToOverwriteSize != -1 {
				var sections []section
				if fileSize == filePercentageToOverwriteSize || _percentage == 100 {
					sections = GetSection(fileSize)
				} else {
					sections = GetSections(fileSize, filePercentageToOverwriteSize)
				}

				if sections != nil && len(sections) > 0 {
					overwriteError = WriteToFile(_file, sections)
				} else {
					fmt.Println(UI_SectionsError)
					overwriteError = true
				}
			}
		} else {
			fmt.Println(UI_FileNotFound, _file)
			overwriteError = true
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_percentage:", _percentage, "_file:", _file)
		overwriteError = true
	}

	return overwriteError
}

// Writes random data to the given file using the given file sections
func WriteToFile(_file string, _sections []section) bool {

	writeError := false

	if len(_file) > 0 && _sections != nil && len(_sections) > 0 {

		var randomData []byte
		var writeBuffer *bufio.Writer

		fileHandle, err := os.OpenFile(_file, os.O_WRONLY, 0)

		if err == nil {
			for counter := 0; counter != len(_sections); counter++ {
				fmt.Print(".")
				if _sections[counter].sectionType == RANDOM_SPACE {
					fileHandle.Seek(_sections[counter].end, 0)
				} else if _sections[counter].sectionType == OVERWRITE_SECTION {
					randomData = GetRandomData(_sections[counter].size)
					if int64(len(randomData)) != 0 && int64(len(randomData)) == _sections[counter].size {

						writeBuffer = bufio.NewWriter(fileHandle)
						writeBuffer = bufio.NewWriterSize(writeBuffer, len(randomData))

						bytesWritten, err := writeBuffer.Write(randomData)

						if err != nil || bytesWritten != len(randomData) {
							fmt.Println(UI_OverwriteError, _file, err.Error())
							writeError = true
							break
						} else {
							// written ok
						}

						writeBuffer.Flush()

					}
				}
			} // end loop

			fileHandle.Sync()
			fileHandle.Close()

		} else {
			fmt.Println(UI_OpenFileError, _file, err.Error())
			writeError = true
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file, "_sections:", _sections)
		writeError = true
	}

	return writeError
}

// Delete the given file
func DeleteFile(_file string) bool {

	deleteError := false

	if len(_file) > 0 {
		err := os.Remove(_file)
		if err != nil {
			fmt.Println(UI_DeleteError, _file, err.Error())
			deleteError = true
		} else {
			// File deleted
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
		deleteError = true
	}

	return deleteError
}
