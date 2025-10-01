package src

import "encoding/binary"

// BYTE_ORDER determines how all binary is en-/decoded in the programme.
var BYTE_ORDER binary.ByteOrder = binary.BigEndian

// VERSION determines the semantic version number output to the token-set header.
var VERSION = [3]byte{0, 0, 1}

var SENTINEL = [4]byte{'L', 'X', 'G', 'O'}
