package main

import (
	"fmt"
	"os/user"
	"runtime"
)

// Returns true if psrm's user has the access rights to write to the given file, false otherwise
func FileUserAccess(_file string) bool {

	fileUserAccess := false

	if len(_file) > 0 {

		switch PLATFORM {
		case PI_ARM_PLATFORM:
			fileUserAccess = FileUserAccessLinux(_file)
		case LINUX_PLATFORM:
			fileUserAccess = FileUserAccessLinux(_file)
		case WINDOWS_PLATFORM:
			fileUserAccess = true
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return fileUserAccess
}

// Returns true if psrm's user has the access rights to write to the given file, false otherwise - linux platform
func FileUserAccessLinux(_file string) bool {

	fileUserAccess := false

	if len(_file) > 0 {

		// get username psrm is running under
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println(UI_GetUserError, _file, err.Error())
		} else {
			// get user of given file
			fileUser, fileUserFound := FileUser(_file)
			if fileUserFound == true {
				if fileUser == currentUser.Username || currentUser.Username == "root" {
					fileUserAccess = true
				}
			}
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
	}

	return fileUserAccess
}

// Gets the name of the currently running function
func GetFunctionName() string {

	data := make([]uintptr, 1)

	var functionName string
	functionName = ""

	callers := runtime.Callers(2, data)
	if callers != 0 {
		caller := runtime.FuncForPC(data[0] - 1)
		if caller != nil {
			functionName = caller.Name()
		}
	}
	return functionName
}
