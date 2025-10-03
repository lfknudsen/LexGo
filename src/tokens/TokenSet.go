package tokens

import (
	"io"
	"os"

	"LexGo/src/config"
)

type TokenSet struct {
	Header TokenSetHeader
	Tokens []Token
}

func NewTokenSet(tokens *[]Token, filename string) *TokenSet {
	result := TokenSet{}
	result.Tokens = *tokens
	name := []byte(filename)
	result.Header = TokenSetHeader{
		Version:        config.VERSION,
		Filename:       name,
		FilenameLength: uint16(len(name)),
		TokenCount:     uint32(len(*tokens)),
	}
	return &result
}

func (ts *TokenSet) Write(w io.Writer) (totalWritten int) {
	totalWritten = ts.Header.Write(w)
	for _, t := range ts.Tokens {
		totalWritten += t.Write(w)
	}
	return totalWritten
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

func DecompileTokenSet(r io.Reader) TokenSet {
	header := DecompileTokenSetHeader(r)

	var tokens = make([]Token, header.TokenCount)
	length := int(header.TokenCount)
	for i := 0; i < length; i++ {
		tokens[i] = DecompileToken(r)
	}
	return TokenSet{header, tokens}
}
