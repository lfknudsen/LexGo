package bin

import (
	"encoding/binary"
	"io"
	"log"
	"unsafe"

	"LexGo/src"
)

type TokenSetHeader struct {
	Version        [3]byte
	TokenCount     uint32
	FilenameLength uint16
	Filename       []rune
}

func (h *TokenSetHeader) Write(w io.Writer) (bytesWritten int) {
	var buffer []byte = make([]byte,
		unsafe.Sizeof(h.Version)+
			unsafe.Sizeof(h.TokenCount)+
			unsafe.Sizeof(h.FilenameLength))
	log.Printf("Allocated a buffer of %d bytes for version, token count,"+
		" and filename length of a tokenset header.\n", len(buffer))
	n1, err := binary.Encode(buffer, src.BYTE_ORDER, h.Version)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Encoded %d bytes for the token set header version.\n", n1)
	n2, err := binary.Encode(buffer[n1:], src.BYTE_ORDER, h.TokenCount)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Encoded %d bytes for the token set header token count.\n", n2)
	n3, err := binary.Encode(buffer[n2:], src.BYTE_ORDER, h.FilenameLength)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Encoded %d bytes for the token set header filename length.\n", n3)
	n4, err := w.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes for the token set header excluding filename itself.\n", n4)
	n := WriteRuneArray(w, h.Filename)
	log.Printf("Wrote %d bytes for the token set filename.\n", n)
	return n + n4
}
