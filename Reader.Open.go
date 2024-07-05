package bitfile

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"os"
)

// Open - Open the named file (error if not exists)
func (o *Reader) Open(fileName *string, blockSize uint) (err error) {

	if blockSize < MinimumBlockSize {

		return fmt.Errorf(errors.ValueTooSmall)

	}

	o.blockSize = blockSize

	if o.file, err = os.Open(*fileName); err != nil {

		err = fmt.Errorf(errors.CannotOpenFile+errors.Details+errors.Details, *fileName, err)

	}

	return err

}
