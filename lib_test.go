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

// Verify and measure the performance of the public API of the gotty package.
package gotty_test

import (
	"io/fs"
	"testing"

	"github.com/go-essentials/gotty"
)

// A "fake" os.File, suitable for testing.
type tOsFile struct {
	// The fs.FileInfo to return when the Stat function is executed.
	statFileInfo gotty.FsFileInfoWrapper

	// The error to return when the Stat function is executed.
	statErr error
}

// A "fake" fs.FileInfo, suitable for testing.
type tFsFileInfo struct {
	// The fs.FileMode to return when the Mode function is executed.
	modeFileMode fs.FileMode
}

// Stat is a "fake" implementation of the os.File.Stat() function.
func (file tOsFile) Stat() (gotty.FsFileInfoWrapper, error) {
	return file.statFileInfo, file.statErr
}

// Mode is a "fake" implementation of the fs.FileInfo.Mode function.
func (fInfo tFsFileInfo) Mode() fs.FileMode {
	return fInfo.modeFileMode
}

// Verify that the IsTTY function is implemented correctly.
func TestIsTTYFalse(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// DEFINITIONS.
	scenarios := []struct {
		name string
		file tOsFile
	}{
		{
			name: "When the `Stdout.Stat()` returns an error.",
			file: tOsFile{statErr: fs.ErrClosed},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeDir`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeDir},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeAppend`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeAppend},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeExclusive`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeExclusive},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeTemporary`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeTemporary},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSymlink`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSymlink},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeDevice`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeDevice},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeNamedPipe`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeNamedPipe},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSocket`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSocket},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSetuid`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSetuid},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSetgid`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSetgid},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSticky`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSticky},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeIrregular`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeIrregular},
			},
		},
	}

	// EXECUTIONS.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		// EXECUTION.
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			gotty.Stdout = scenario.file

			// ACT / ASSERT.
			if gotty.IsTTY() {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected: false\033[0m\n"+
					"\033[31mActual:   true\033[0m\n\n",
					scenario.name)
			}
		})
	}
}

// Verify that the IsTTY function is implemented correctly.
func TestIsTTYTrue(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// DEFINITIONS.
	scenarios := []struct {
		name string
		file tOsFile
	}{
		{
			name: "When the `Stdout` represents a `fs.ModeCharDevice`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeCharDevice},
			},
		},
	}

	// EXECUTIONS.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		// EXECUTION.
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			gotty.Stdout = scenario.file

			// ACT / ASSERT.
			if !gotty.IsTTY() {
				t.Fatalf("\n\n"+
					"UT Name:  %s\n"+
					"\033[32mExpected: true\033[0m\n"+
					"\033[31mActual:   false\033[0m\n\n",
					scenario.name)
			}
		})
	}
}

// Measure the performance of the IsTTY function.
func BenchmarkIsTTY(b *testing.B) {
	// DEFINITIONS.
	scenarios := []struct {
		name string
		file tOsFile
	}{
		{
			name: "When the `Stdout.Stat()` returns an error.",
			file: tOsFile{statErr: fs.ErrClosed},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeDir`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeDir},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeAppend`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeAppend},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeExclusive`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeExclusive},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeTemporary`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeTemporary},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSymlink`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSymlink},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeDevice`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeDevice},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeNamedPipe`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeNamedPipe},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSocket`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSocket},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSetuid`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSetuid},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSetgid`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSetgid},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeCharDevice`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeCharDevice},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeSticky`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeSticky},
			},
		},
		{
			name: "When the `Stdout` represents a `fs.ModeIrregular`.",
			file: tOsFile{
				statFileInfo: tFsFileInfo{modeFileMode: fs.ModeIrregular},
			},
		},
	}

	// EXECUTION.
	for _, scenario := range scenarios {
		scenario := scenario // NOTE: Ensure that the b.Run function has the correct value when it's being executed.

		b.Run(scenario.name, func(b *testing.B) {
			// ARRANGE.
			gotty.Stdout = scenario.file

			// RESET.
			b.ResetTimer()

			// RUN.
			for n := 0; n < b.N; n++ {
				// ACT.
				_ = gotty.IsTTY()
			}
		})
	}
}
