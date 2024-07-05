package bitfile

import (
	"os"
	"testing"
)

func TestBitFile_Name(t *testing.T) {
	const (
		testFileName = "Bitfile.Name.test.txt"
	)
	var (
		bitfile Reader
		err     error
	)

	if bitfile.Name() != "" {
		t.Fatal("expect empty string when file not open")
	}

	bitfile.file, err = os.CreateTemp("", testFileName)

	if err != nil {
		t.Fatal(err)
	}

	if bitfile.Name() == "" {
		t.Fatal("unexpected empty file name")
	}

	if bitfile.file.Name() != bitfile.Name() {
		t.Fatal("filename mismatch")
	}
}
