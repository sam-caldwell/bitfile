package bitfile

import (
	"os"
	"testing"
)

// TestBitFile_Close - create a test file, test the Reader.Close() method and then clean up afterward.
func TestBitFile_Close(t *testing.T) {
	const (
		testFileName = "/tmp/TestBitFile_Close.txt"
		testData     = "this is a test"
	)
	var testFile string

	t.Cleanup(func() {
		if err := os.Remove(testFileName); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("create test file", func(t *testing.T) {
		var (
			err error
			f   *os.File
		)
		if f, err = os.Create(testFileName); err != nil {
			t.Fatal(err)
		}
		testFile = f.Name()
		n, err := f.Write([]byte(testData))
		if err != nil {
			t.Fatal("error writing test data")
		}
		if n != len([]byte(testData)) {
			t.Fatal("test data length mismatch.")
		}
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Perform a .Close() with a nil file handle.  Expect no error", func(t *testing.T) {
		var f Reader
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Perform the test of the bitfile.Close() method.", func(t *testing.T) {
		var f Reader
		if err := f.Open(&testFile, MinimumBlockSize); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				t.Fatal(err)
			}
		}()
	})

	t.Run("Perform the test of the bitfile.Close() method with double-close", func(t *testing.T) {
		var f Reader
		if err := f.Open(&testFile, MinimumBlockSize); err != nil {
			t.Fatal(err)
		}
		defer func() {
			var f Reader
			if err := f.Close(); err != nil {
				t.Fatal(err)
			}
		}()
	})
}
