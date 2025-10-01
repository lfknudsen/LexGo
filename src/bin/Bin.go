package bin

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"unsafe"

	"LexGo/src"
)

type File struct {
	Header  FileHeader
	Content FileContent
}

func NewBinFile(tokenSet ...TokenSet) *File {
	result := File{}
	result.Header = FileHeader{
		Sentinel:         src.SENTINEL,
		Version:          src.VERSION,
		TokenSetCount:    uint16(len(tokenSet)),
		TokenSetHeaderSz: uint8(unsafe.Sizeof(TokenSetHeader{})),
	}
	result.Content = tokenSet
	log.Printf("Created binary file structure.\n")
	return &result
}

func (b *File) Write(w io.Writer) (totalWritten int) {
	headerN := b.Header.Write(w)
	log.Printf("Wrote binary file header to disk; %d bytes.\n", headerN)
	return headerN + b.Content.Write(w)
}

func Write(tokenSet *TokenSet, filename string) {
	outputFile, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Created output file: %v.\n", filename)
	// writer := bufio.NewWriter(outputFile)
	bin := NewBinFile(*tokenSet)
	n := bin.Write(outputFile)
	log.Printf("Finished writing binary file to disk; %d bytes.\n", n)

	defer func(outputFile *os.File) {
		err := outputFile.Sync()
		if err != nil {
			log.Panic(err)
		}
		err = outputFile.Close()
		log.Printf("Closed file %v.\n", outputFile.Name())
		if err != nil {
			log.Panic(err)
		}
	}(outputFile)
}

func EncodeRuneArray(runes []rune) []byte {
	flatSize := 0
	for i := 0; i < len(runes); i++ {
		flatSize += binary.Size(runes[i])
	}

	output := make([]byte, flatSize)
	offset := 0
	for i := 0; i < len(runes); i++ {
		n, err := binary.Encode(output[offset:], src.BYTE_ORDER, runes[i])
		if err != nil {
			log.Panic(err)
		}
		offset += n
	}
	return output
}

func WriteRuneArray(w io.Writer, runes []rune) (bytesWritten int) {
	buffer := make([]byte, 4) // rune size
	totalWritten := 0
	for i := 0; i < len(runes); i++ {
		n, err := binary.Encode(buffer, src.BYTE_ORDER, runes[i])
		if err != nil {
			log.Panic(err)
		}
		n, err = w.Write(buffer[0:n])
		if err != nil {
			log.Panic(err)
		}
		totalWritten += n
	}
	return totalWritten
}

func DecompileBinFile(r io.Reader) *File {
	header := DecompileBinHeader(r)
	content := DecompileBinContent(r)
	output := File{*header, *content}
	return &output
}
