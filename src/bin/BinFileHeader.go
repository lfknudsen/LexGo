package bin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/lfknudsen/golib/src/maths"
	"github.com/lfknudsen/golib/src/structs"

	"LexGo/src"
)

type FileHeader struct {
	Sentinel         [4]byte
	Version          structs.Version
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

func (h *FileHeader) Print() {
	h.PrintTo(os.Stdout)
}

func (h *FileHeader) PrintTo(w io.Writer) {
	line1 := fmt.Sprintf("# Sentinel: %s | Version: %s",
		h.Sentinel, h.Version.String())
	line2 := fmt.Sprintf("# Token sets: %d | Token set header size (bytes): %d\n",
		h.TokenSetCount, h.TokenSetHeaderSz)
	borderLength := max(len(line1), len(line2)) + 2
	lineDifference := maths.Abs(len(line1) - len(line2))
	padding := strings.Repeat(" ", lineDifference+1)
	line1 += padding
	line1 += " #"
	line2 += " #"
	border := bytes.Repeat([]byte{'#'}, borderLength)
	_, _ = w.Write(border)
	_, _ = w.Write([]byte(line1))
	_, _ = w.Write([]byte(line2))
	_, _ = w.Write(border)
}
