package main

const (
	UI_Help = `psrm v1.0 by gburnett@outlook.com

psrm overwrites the given percentage of a file with random data - sections of file to overwrite are randomly chosen. File is then deleted.

Syntax: 
	psrm percentage file
	psrm percentage file1 file2
	psrm percentage *

Arguments:
	percentage - percentage of file to overwrite
	file - file to overwrite - single file name, multiple files, or wildcard
	
Examples:
	psrm 10 MyFile - overwrite 10 percent of MyFile, then delete it
	psrm 30 MyFile1 MyFile2 - overwrite 30 percent of MyFile1 and MyFile2, then delete the files
	psrm 60 * - overwrite 60 percent of all files in current directory

Website:`
	UI_InvalidArgs      = `ERROR: Invalid arguments`
	UI_InvalidPercent   = `ERROR: Invalid percentage`
	UI_FileStatError    = `ERROR: Failed to get file stats`
	UI_GetUserError     = `ERROR: Unable to get current user`
	UI_FileUserError    = `ERROR: Unable to get file user`
	UI_FileNoUserPerm   = `ERROR: No user permission:`
	UI_FileNoWritePerm  = `ERROR: No write permission:`
	UI_FileNotFound     = `ERROR: File not found:`
	UI_EmptyFile        = `ERROR: Empty file`
	UI_InvalidFileSize  = `ERROR: Invalid file size`
	UI_FileTooBig       = `ERROR: File too big`
	UI_NoFileSize       = `ERROR: Can't get file size`
	UI_ParameterInvalid = `ERROR: Invalid parameter: %s , Parameters: `
	UI_DeleteError      = `ERROR: Delete error: `
	UI_OverwriteError   = `ERROR: Overwrite error: `
	UI_SectionsError    = `ERROR: No file sections`
	UI_OpenFileError    = `ERROR: Open file failed: `
	UI_RandomDataError  = `ERROR: Random data not generated`

	UI_Deleted          = `Deleted`
	UI_RandomSpace      = `Random Space`
	UI_OverwriteSection = `Overwrite Section`
)
