package bin

import (
	"io"
	"os"

	"LexGo/src/tokens"
)

type FileContent struct {
	TokenSets []tokens.TokenSet
}

func NewFileContent(tokenSets []tokens.TokenSet) FileContent {
	return FileContent{TokenSets: tokenSets}
}

func (c *FileContent) Write(w io.Writer) (totalWritten int) {
	totalWritten = 0
	length := len(c.TokenSets)
	for i := 0; i < length; i++ {
		totalWritten += c.TokenSets[i].Write(w)
	}
	return totalWritten
}

func DecompileBinContent(r io.Reader, header FileHeader) FileContent {
	content := FileContent{}
	content.TokenSets = make([]tokens.TokenSet, header.TokenSetCount)
	length := len(content.TokenSets)
	for i := 0; i < length; i++ {
		content.TokenSets[i] = tokens.DecompileTokenSet(r)
	}
	return content
}

func (c *FileContent) Print() {
	c.PrintTo(os.Stdout)
}

func (c *FileContent) PrintTo(out io.Writer) {
	for _, tokenSet := range c.TokenSets {
		tokenSet.PrintTo(out)
	}
}
