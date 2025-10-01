package src

import (
	"log"
	"os"
	"regexp"
	"strings"
)

// LoadRegex loads a regex string from a plaintext file. It skips over whitespace, meaning
// you are free to structure and indent the file however you wish.
func LoadRegex(filename string) *Regex {
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
	//log.Println(sb.String())
	re := regexp.MustCompile(sb.String())
	return NewRegex(re)
}
