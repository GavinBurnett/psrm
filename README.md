# psrm

psrm - Partial secure remove - the fast file shredder.

By: gburnett@outlook.com

This program overwrites the given file(s) with random data. The percentage of the file to overwrite is set by the user. The whole file is then deleted.

Useful for speeding up the destruction of large files that do not need to be entirely destroyed, but a standard delete is not secure enough.

This software is used at the users own risk.
The author takes no responsibility for data loss.

Released under the GPL 3 licence.

Installation:

Copy the executable relevant to the platform to a directory that is included in the path search for the command line.

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

Error codes returned to shell:

0  = success
-1 = failed
