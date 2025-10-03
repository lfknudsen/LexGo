package bin

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"unsafe"

	"github.com/lfknudsen/golib/src/structs"

	"LexGo/src"
)

type TokenSetHeader struct {
	Version        structs.Version
	TokenCount     uint32
	FilenameLength uint16 // in bytes
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

func (h *TokenSetHeader) Print() {
	h.PrintTo(os.Stdout)
}

func (h *TokenSetHeader) PrintTo(out io.Writer) {
	_, _ = fmt.Fprintf(out, "# Version: %s | Tokens: %d\n", h.Version.String(), h.TokenCount)
	_, _ = fmt.Fprintf(out, "# Source: %s\n", string(h.Filename))
}

func DecompileTokenSetHeader(r io.Reader, TokenSetHeaderSz uint8) *TokenSetHeader {
	version, err := structs.DecompileVersion(r)
	if err != nil {
		log.Panic(err)
	}
	var tokenCount uint32
	buffer := make([]byte, unsafe.Sizeof(tokenCount))
	_, err = r.Read(buffer)
	if err != nil {
		log.Panic(err)
	}
	_, err2 := binary.Decode(buffer, src.BYTE_ORDER, &tokenCount)
	if err2 != nil {
		log.Panic(err2)
	}

	var filenameLength uint16
	buffer = make([]byte, unsafe.Sizeof(filenameLength))
	_, err = r.Read(buffer)
	if err != nil {
		log.Panic(err)
	}
	_, err = binary.Decode(buffer, src.BYTE_ORDER, &filenameLength)
	if err != nil {
		log.Panic(err)
	}

	buffer = make([]byte, filenameLength)
	filename := make([]rune, filenameLength)
	_, err = r.Read(buffer)
	if err != nil {
		log.Panic(err)
	}
	_, err = binary.Decode(buffer, src.BYTE_ORDER, &filename)
	if err != nil {
		log.Panic(err)
	}
	return &TokenSetHeader{
		Version:        *version,
		TokenCount:     tokenCount,
		FilenameLength: filenameLength,
		Filename:       filename,
	}
}
