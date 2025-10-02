package src

import (
	"encoding/binary"

	"github.com/lfknudsen/golib/src/structs"
)

// BYTE_ORDER determines how all binary is en-/decoded in the programme.
var BYTE_ORDER binary.ByteOrder = binary.BigEndian

// WRITE_BOM determines if the output file should begin with a Byte Order Mark.
const WRITE_BOM = true

// VERSION determines the semantic version number output to the token-set header.
var VERSION = structs.Version{Major: 0, Minor: 0, Patch: 1}

var SENTINEL = [4]byte{'L', 'X', 'G', 'O'}
