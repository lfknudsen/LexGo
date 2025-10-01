package src

import (
	"encoding/binary"
	"io"
	"log"
	"strings"
)

func RuneToString(rs []rune) string {
	sb := new(strings.Builder)
	sb.Grow(len(rs))
	for _, r := range rs {
		_, _ = sb.WriteRune(r)
	}
	return sb.String()
}

func ArraysEqual(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func WriteByteArray(w io.Writer, in []byte) {
	err := binary.Write(w, binary.BigEndian, int32(len(in)))
	if err != nil {
		log.Panic(err)
	}
	for i := 0; i < len(in); i++ {
		err = binary.Write(w, binary.BigEndian, in[i])
		if err != nil {
			log.Panic(err)
		}
	}
}

func ReadByteArray(r io.Reader) []byte {
	var sz int32
	err := binary.Read(r, binary.BigEndian, &sz)
	if err != nil {
		log.Panic(err)
	}

	array := make([]byte, sz)
	err = binary.Read(r, binary.BigEndian, &array)
	if err != nil {
		log.Panic(err)
	}
	return array
}
