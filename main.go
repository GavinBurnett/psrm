// psrm project main.go
package main

import (
	"fmt"
	"os"
)

func main() {

	exitCode := 0

	if os.Args != nil {

		args := os.Args
		validArgs := false

		percent := 0
		files := make([]string, len(args))

		// args[0] = psrm
		// args[1] = percentage
		// args[2,3,4,..] = file / files

		if len(args) == 1 {
			// no user arguments given - display help
			fmt.Println(UI_Help)
		}
		if len(args) == 2 {
			if IsStringHelpArgument(args[1]) {
				// user has given help argument - display help
				fmt.Println(UI_Help)
			} else {
				// user has given only one argument that is not a help argument - display error
				exitCode = -1
				fmt.Println(UI_InvalidArgs)
			}
		}
		if len(args) == 3 {
			// user has given percentage and one file name arguments

			// store file in array
			files = args[2:len(args)]

			validArgs = true
		}
		if len(args) > 3 {
			// user has given percentage and multiple files arguments

			// store all files in array
			files = args[2:len(args)]

			validArgs = true
		}

		if validArgs == true {
			if IsStringNumber(args[1]) {
				fmt.Sscan(args[1], &percent)
				if IsPercentageValid(percent) {
					if ProcessFiles(percent, files) {
						// All done with no errors
					} else {
						// All done with errors
						exitCode = -1
					}
				} else {
					exitCode = -1
					fmt.Println(UI_InvalidPercent)
				}
			} else {
				exitCode = -1
				fmt.Println(UI_InvalidPercent)
			}
		}
	} else {
		exitCode = -1
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
	}

	os.Exit(exitCode)
}
