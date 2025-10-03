package bin

import (
	"io"
	"log"
	"os"

	"LexGo/src"
)

type TokenSet struct {
	Header TokenSetHeader
	Tokens []src.Token
}

func NewTokenSet(tokens *[]src.Token, filename string) *TokenSet {
	result := TokenSet{}
	result.Tokens = *tokens
	name := []rune(filename)
	result.Header = TokenSetHeader{
		Version:        src.VERSION,
		Filename:       name,
		FilenameLength: uint16(len(name)),
		TokenCount:     uint32(len(*tokens)),
	}
	return &result
}

func (ts *TokenSet) Write(w io.Writer) (totalWritten int) {
	headerN := ts.Header.Write(w)
	log.Printf("Finished writing tokenset header. Wrote %d bytes.\n", headerN)
	totalWritten = 0
	for _, t := range ts.Tokens {
		n := t.Write(w)
		totalWritten += n
	}
	log.Printf("Finished writing tokenset tokens. Wrote %d bytes (excl. header).\n",
		totalWritten)
	return headerN + totalWritten
}

func (ts *TokenSet) Print() {
	ts.PrintTo(os.Stdout)
}

func (ts *TokenSet) PrintTo(out io.Writer) {
	ts.Header.PrintTo(out)
	for _, token := range ts.Tokens {
		token.PrintTo(out)
	}
}

func DecompileTokenSet(r io.Reader, TokenSetHeaderSz uint8) *TokenSet {
	header := DecompileTokenSetHeader(r, TokenSetHeaderSz)
	var tokens []src.Token
	for i := 0; i < int(header.TokenCount); i++ {
		tokens = append(tokens, *src.DecompileToken(r))
	}
	return &TokenSet{*header, tokens}
}
