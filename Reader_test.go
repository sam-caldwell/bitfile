package bitfile

import (
	"github.com/sam-caldwell/size"
	"testing"
)

// TestBitFile_Struct - test the structure itself.
//
//goland:noinspection ALL
func TestBitFile_Struct(t *testing.T) {
	var b Reader

	//Validate the constant used for default buffer size
	if defaultBufferSize < 4*size.MB {
		t.Fatal("Bitfile should have a constant defaultBufferSize of at least 1024")
	}

	if b.file != nil {
		t.Fatal("Bitfile.file should be nil initially.")
	}
}
