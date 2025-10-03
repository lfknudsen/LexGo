package bin

import (
	"io"
	"log"
	"os"

	"LexGo/src/config"
	"LexGo/src/tokens"
)

type File struct {
	BOM     BOM
	Header  FileHeader
	Content FileContent
}

func NewBinFile(tokenSets []tokens.TokenSet) *File {
	result := File{}
	result.BOM = NewBOM()
	result.Header = NewFileHeader(config.SENTINEL, config.VERSION, int32(len(tokenSets)))
	result.Content = NewFileContent(tokenSets)
	log.Printf("Created binary file structure.\n")
	return &result
}

func (b *File) Write(w io.Writer) (totalWritten int) {
	n := b.BOM.Write(w)
	log.Printf("Wrote %d bytes for BOM.\n", n)
	headerN := b.Header.Write(w)
	log.Printf("Wrote binary file header to disk; %d bytes.\n", headerN)
	contentN := b.Content.Write(w)
	log.Printf("Wrote binary file content to disk; %d bytes.\n", contentN)
	return n + headerN + contentN
}

func Write(tokenSets []tokens.TokenSet, filename string) {
	outputFile, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Created output file: %v.\n", filename)

	bin := NewBinFile(tokenSets)

	bin.Print()

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

func (b *File) Print() {
	b.Header.Print()
	b.Content.Print()
}

func (b *File) PrintTo(out io.Writer) {
	b.Header.PrintTo(out)
	b.Content.PrintTo(out)
}

func DecompileBinFile(r io.Reader) *File {
	BOM := DecompileBOM(r)
	header := DecompileBinHeader(r)
	header.Print()
	content := DecompileBinContent(r, header)
	output := File{BOM, header, content}
	return &output
}
