package bitfile

import (
	"os"
)

// Reader - A Bitfile struct represents a single bitwise reader object
type Reader struct {
	blockSize uint
	// file handle for the associated file
	file *os.File
}
