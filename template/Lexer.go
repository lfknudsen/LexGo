package template

import (
	"fmt"
	"log"
	"os"
	"regexp"

	. "LexGo/src"
	. "LexGo/src/bin"
	. "LexGo/src/regex"
	. "LexGo/src/tokens"
)

const Name = "LexGo"

func OpenCodeFile(filename string) (outputFilename string) {
	code, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	regexFile, err := os.ReadFile("in_re.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("About to compile...\n")
	re := regexp.MustCompile(string(regexFile))
	rex := NewRegex(re)
	tokens := LexTokens(rex, &code)
	tokenset := NewTokenSet(tokens, filename)
	outputFilename = filename + "_out.txt"
	Write([]TokenSet{*tokenset}, outputFilename)
	return outputFilename
}

// ReadTokens generates an array of src.Token structs from the given regular expression,
// matched against the given byte-array. Convenience function which calls [LexTokens].
func ReadTokens(regexp *regexp.Regexp, code *[]byte) *[]Token {
	return LexTokens(NewRegex(regexp), code)
}

// LexTokens generates an array of src.Token structs from the given regular expression,
// matched against the given byte-array.
func LexTokens(regex *Regex, code *[]byte) *[]Token {
	ruleset := Decompile(regex.Src())
	if ruleset == nil {
		log.Panic("Ruleset is nil")
	}
	names := ruleset.Names()

	fmt.Println("Subexpression names:")
	for idx, name := range *names {
		fmt.Printf("%d: %s\n", idx, name)
	}

	var tokenIDs = make([]string, 0)
	var values = make([]string, 0)
	tokens := make([]Token, 0)

	fmt.Println("Beginning matching.")
	matches := regex.FindAllSubmatchIndex(code)
	for _, match := range matches {
		for i, rule := range ruleset.Rules {
			if rule.Id == "" {
				continue
			}
			idx := regex.SubNames[rule.Id]
			if idx == -1 {
				continue
			}
			left, right := match[idx*2], match[idx*2+1]
			if left != -1 && left != right {
				value := (*code)[left:right]
				token := Token{
					ID:    byte(i),
					Value: make([]byte, 0),
				}
				token.Value = value
				token.ValueLength = uint16(len(value))

				values = append(values, string(value))
				tokenIDs = append(tokenIDs, rule.Id)
				tokens = append(tokens, token)
				break
			}
		}
	}
	fmt.Println("Result of pattern matching on code text:")
	for i, value := range values {
		fmt.Println(tokenIDs[i] + ":    " + string(value))
	}
	return &tokens
}
