package tokens

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lfknudsen/golib/src/structs"

	"LexGo/src/config"
)

type TokenSetHeader struct {
	TokenCount     uint32
	FilenameLength uint16 // in bytes
	Filename       []byte
}

func (h *TokenSetHeader) Write(w io.Writer) (bytesWritten int) {
	err := binary.Write(w, config.BYTE_ORDER, h.TokenCount)
	if err != nil {
		log.Fatal(err)
	}
	err = binary.Write(w, config.BYTE_ORDER, h.FilenameLength)
	if err != nil {
		log.Fatal(err)
	}
	err = binary.Write(w, config.BYTE_ORDER, h.Filename)
	if err != nil {
		log.Fatal(err)
	}
	return binary.Size(h.TokenCount) +
		binary.Size(h.FilenameLength) + binary.Size(h.Filename)
}

func (h *TokenSetHeader) Print() {
	h.PrintTo(os.Stdout)
}

func (h *TokenSetHeader) PrintTo(out io.Writer) {
	_, err := fmt.Fprintf(out, "+ Tokens: %d | Filename length: %d\n", h.TokenCount, h.FilenameLength)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(out, "+ Source: %s\n", string(h.Filename))
	if err != nil {
		log.Fatal(err)
	}
}

func DecompileTokenSetHeader(r io.Reader, version structs.Version) TokenSetHeader {
	ts := TokenSetHeader{}

	if version.IsLowerThan(1, 1, 0) {
		_, err := structs.DecompileVersion(r)
		if err != nil {
			log.Panic(err)
		}
	}

	err := binary.Read(r, config.BYTE_ORDER, &ts.TokenCount)
	if err != nil {
		log.Panic(err)
	}

	err = binary.Read(r, config.BYTE_ORDER, &ts.FilenameLength)
	if err != nil {
		log.Panic(err)
	}
	ts.Filename = make([]byte, ts.FilenameLength)
	err = binary.Read(r, config.BYTE_ORDER, &ts.Filename)
	if err != nil {
		log.Panic(err)
	}
	return ts
}
