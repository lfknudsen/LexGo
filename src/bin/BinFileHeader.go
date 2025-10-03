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

	"LexGo/src/config"
)

type FileHeader struct {
	Sentinel      config.Sentinel
	Version       structs.Version
	TokenSetCount int32
}

func NewFileHeader[T maths.Signed](sentinel config.Sentinel, version structs.Version, tokenCount T) FileHeader {
	return FileHeader{
		Sentinel:      sentinel,
		Version:       version,
		TokenSetCount: int32(tokenCount),
	}
}

func (h *FileHeader) Write(w io.Writer) (totalWritten int) {
	var buffer = make([]byte, binary.Size(*h))
	n1, err := binary.Encode(buffer, config.BYTE_ORDER, *h)
	if err != nil {
		log.Panic(err)
	}
	n2, err := w.Write(buffer)
	if err != nil {
		log.Panic(err)
	}
	if n1 != n2 {
		log.Panicf("Encoded %d bytes, but wrote %d!\n", n1, n2)
	}
	return n2
}

func DecompileBinHeader(r io.Reader) FileHeader {
	var output FileHeader
	err := binary.Read(r, config.BYTE_ORDER, &output)
	if err != nil {
		log.Panic(err)
	}
	return output
}

func (h *FileHeader) Print() {
	h.PrintTo(os.Stdout)
}

func (h *FileHeader) PrintTo(w io.Writer) {
	line1 := fmt.Sprintf("# Sentinel: %s | Version: %s",
		h.Sentinel.String(), h.Version.String())
	line2 := fmt.Sprintf("# Token sets:  %d",
		h.TokenSetCount)
	borderLength := max(len(line1), len(line2)) + 2
	lineDifference := maths.Abs(len(line1) - len(line2))
	padding := strings.Repeat(" ", lineDifference)
	line2 += padding
	line1 += " #\n"
	line2 += " #\n"
	border := bytes.Repeat([]byte{'#'}, borderLength+1)
	border[borderLength] = byte('\n')
	_, _ = w.Write(border)
	_, _ = w.Write([]byte(line1))
	_, _ = w.Write([]byte(line2))
	_, _ = w.Write(border)
}
