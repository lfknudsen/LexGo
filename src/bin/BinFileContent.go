package bin

import (
	"io"
	"log"
	"os"
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

func DecompileBinContent(r io.Reader, header FileHeader) *FileContent {
	output := make(FileContent, header.TokenSetCount)
	for i := 0; i < len(output); i++ {
		output[i] = *DecompileTokenSet(r, header.TokenSetHeaderSz)
	}
	return &output
}

func (c *FileContent) Print() {
	c.PrintTo(os.Stdout)
}

func (c *FileContent) PrintTo(out io.Writer) {
	for _, tokenSet := range *c {
		tokenSet.PrintTo(out)
	}
}
