package bitfile

import (
	"fmt"
	"os"
)

// Size - Return file size (in bytes)
func (o *Writer) Size() (sz uint64, err error) {

	if o.file == nil {
		return 0, fmt.Errorf("file unknown")
	}

	fileInfo, err := os.Stat(o.file.Name())

	if err != nil {
		return 0, err
	}

	return uint64(fileInfo.Size()), nil

}
