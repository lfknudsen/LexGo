package bin

import (
	"encoding/binary"
	"io"
	"log"

	"LexGo/src"
)

type FileHeader struct {
	Sentinel         [4]byte
	Version          [3]byte
	TokenSetCount    uint16
	TokenSetHeaderSz uint8
}

func (h *FileHeader) Write(w io.Writer) (totalWritten int) {
	var buffer []byte = make([]byte, binary.Size(*h))
	log.Printf("Allocated buffer for binary file header of %d bytes.\n",
		len(buffer))
	n, err := binary.Encode(buffer, src.BYTE_ORDER, *h)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Encoded binary file header; %d bytes.\n", n)
	n, err = w.Write(buffer)
	if err != nil {
		log.Panic(err)
	}
	return n
}

func DecompileBinHeader(r io.Reader) *FileHeader {
	var output FileHeader
	buffer := make([]byte, binary.Size(FileHeader{}))
	_, err := r.Read(buffer)
	if err != nil {
		log.Panic(err)
	}
	_, err = binary.Decode(buffer, src.BYTE_ORDER, &output)
	if err != nil {
		log.Panic(err)
	}
	return &output
}
