package config

import (
	"encoding/binary"

	"github.com/lfknudsen/golib/src/structs"
)

// BYTE_ORDER determines how all binary is en-/decoded in the programme.
var BYTE_ORDER binary.ByteOrder = binary.BigEndian

// USE_BOM determines if the output file should begin with a Byte Order Mark, and
// if the programme should expect a BOM when reading such a file back.
var USE_BOM = true

// VERSION determines the semantic version number output to the token-set header.
var VERSION = structs.Version{Major: 0, Minor: 9, Patch: 0}

// SENTINEL is the four bytes which are written first to the binary file (right after the
// byte order mark)
var SENTINEL = TrueSentinel()

func ToggleByteOrder() {
	switch BYTE_ORDER {
	case binary.BigEndian:
		BYTE_ORDER = binary.LittleEndian
	case binary.LittleEndian, binary.NativeEndian:
		BYTE_ORDER = binary.BigEndian
	}
}
