package bitfile

import (
	"github.com/sam-caldwell/bitblock"
)

// Read - Read a block of bytes from the file and return the block to the caller.
func (o *Reader) Read() (block *bitblock.Block, err error) {

	//var bytesRead int // count of bytes read
	var buffer []byte

	block = bitblock.NewBlock(o.blockSize)

	buffer = make([]byte, block.Size())

	if _, err = o.file.Read(buffer); err != nil {
		return nil, err
	}

	block.Set(buffer)

	//if (bytesRead > 0) && (bytesRead < len(buffer)) {
	//	block.Set(buffer)
	//}

	return block, err
}
