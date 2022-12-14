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

// Package gotty implements functions to detect that a Go application is running inside a terminal.
package gotty

import (
	"io/fs"
	"os"
)

// A subset around Go's standard os.File struct.
type osFileWrapper interface {
	Stat() (FsFileInfoWrapper, error)
}

// A subset of Go's standard fs.FileInfo interface.
type FsFileInfoWrapper interface {
	// Returns the file's mode.
	Mode() fs.FileMode
}

// A wrapper around an os.File pointer.
type osFile struct {
	// The wrapped os.File pointer.
	file *os.File
}

// A wrapper around a fs.FileInfo instance.
type fsFileInfo struct {
	// The wrapper fs.FileInfo instance.
	fileInfo fs.FileInfo
}

// Wraps Stdout.
var Stdout osFileWrapper = osFile{file: os.Stdout}

// IsTTY returns true if Stdout is attached to a terminal window, false otherwise.
func IsTTY() bool {
	if fileInfo, err := Stdout.Stat(); err == nil {
		return fileInfo.Mode()&os.ModeCharDevice != 0
	}
	return false
}
