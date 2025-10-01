package bin

import (
	"encoding/binary"
	"io"
	"log"

	"LexGo/src"
)

type FileContent []TokenSet

func (c *FileContent) Write(w io.Writer) (totalWritten int) {
	log.Printf("Writing binary file contents to disk.\n")
	totalWritten = 0
	for i := 0; i < len(*c); i++ {
		n := (*c)[i].Write(w)
		totalWritten += n
	}
	log.Printf("Wrote binary file contents to disk; %d bytes.\n", totalWritten)
	return totalWritten
}

func DecompileBinContent(r io.Reader) *FileContent {
	var output FileContent
	buffer := make([]byte, binary.Size(FileContent{}))
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
