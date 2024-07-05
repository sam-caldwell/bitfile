package bitfile

import (
	"github.com/sam-caldwell/bitblock"
)

// ReadBytes - Read a sequence of n bytes from the current file and return them as a Block.
//
//	Reader.ReadBytes() will read a block of bytes from a given file
//	and return a Block struct from which individual bits can be read
//	or other block operations can be performed.
func (o *Reader) ReadBytes(n uint) (block *bitblock.Block, err error) {

	block = bitblock.NewBlock(n)

	data := make([]byte, n)

	_, err = o.file.Read(data)

	block.Set(data)

	return block, err

}
