// =====================================================================================================================
// = LICENSE:       Copyright (c) 2022 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// This file contains functions which should NOT be excluded from code coverage.
// The functions defined in this class integrated with the underlying OS, and are impossible to test.

// Package gotty implements functions to detect that a Go application is running inside a terminal.
package gotty

import "io/fs"

// Stat returns the FileInfo structure describing file.
// If there is an error, it will be of type *PathError.
func (file osFile) Stat() (FsFileInfoWrapper, error) {
	if fileInfo, err := file.file.Stat(); err != nil {
		return nil, err
	} else {
		return fsFileInfo{fileInfo: fileInfo}, nil
	}
}

// Mode returns the file's mode.
func (fileInfo fsFileInfo) Mode() fs.FileMode {
	return fileInfo.fileInfo.Mode()
}
