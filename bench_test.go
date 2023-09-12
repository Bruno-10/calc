package main

import (
	"os"
	"testing"
)

func Benchmark_firstImpl(b *testing.B) {
	input := "1+1+1;2+2+2\n"
	for i := 0; i < b.N; i++ {
		firstImpl(input) // prints Group 1: 3; Group 2: 6; Total Sum: 9;
	}
}

func Benchmark_secondImpl(b *testing.B) {
	input := "1+1+1;2+2+2\n"
	for i := 0; i < b.N; i++ {
		secondImpl(input) // prints Group 1: 3; Group 2: 6; Total Sum: 9;
	}
}

// ----------------------------------------------------------------------------
//  Helper Functions
// ----------------------------------------------------------------------------

// mockStdin is a helper function that lets the test pretend dummyInput as os.Stdin.
// It will return a function to `defer` clean up after the test.
func mockStdin(b *testing.B, dummyInput string) (funcDefer func(), err error) {
	b.Helper()

	oldOsStdin := os.Stdin

	tmpfile, err := os.CreateTemp(b.TempDir(), b.Name())
	if err != nil {
		return nil, err
	}

	content := []byte(dummyInput)

	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return nil, err
	}

	// Set stdin to the temp file
	os.Stdin = tmpfile

	return func() {
		// clean up
		os.Stdin = oldOsStdin
		os.Remove(tmpfile.Name())
	}, nil
}
