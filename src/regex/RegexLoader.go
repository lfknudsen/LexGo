package regex

import (
	"log"
	"os"
	"regexp"
	"strings"
)

// LoadRegexp loads a regexp string from a plaintext file. It skips over whitespace,
// meaning you are free to structure and indent the file however you wish.
func LoadRegexp(filename string) *regexp.Regexp {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	sb := strings.Builder{}
	sb.Grow(len(fileBytes))
	for _, char := range fileBytes {
		if char != '\n' && char != '\r' && char != '\t' && char != ' ' {
			sb.WriteByte(char)
		}
	}
	return regexp.MustCompile(sb.String())
}

// LoadRegex loads a regexp string from a plaintext file. It skips over whitespace,
// meaning you are free to structure and indent the file however you wish.
func LoadRegex(filename string) *Regex {
	return NewRegex(LoadRegexp(filename))
}
