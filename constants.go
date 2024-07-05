package bitfile

import (
	"github.com/sam-caldwell/size"
)

const (

	//defaultBufferSize - This is the default buffer size (4MB)
	defaultBufferSize = 4 * size.MB

	//MinimumBlockSize - The minimum block size read from / written to a file
	MinimumBlockSize = 64
)
