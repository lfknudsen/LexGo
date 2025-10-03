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
	return &result
}

func (b *File) Write(w io.Writer) (totalWritten int) {
	n := b.BOM.Write(w)
	headerN := b.Header.Write(w)
	contentN := b.Content.Write(w)
	return n + headerN + contentN
}

func Write(tokenSets []tokens.TokenSet, filename string) {
	outputFile, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Sync()
		if err != nil {
			log.Panic(err)
		}
		err = outputFile.Close()
		if err != nil {
			log.Panic(err)
		}
	}(outputFile)

	bin := NewBinFile(tokenSets)
	_ = bin.Write(outputFile)
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
	content := DecompileBinContent(r, header)
	output := File{BOM, header, content}
	return &output
}
