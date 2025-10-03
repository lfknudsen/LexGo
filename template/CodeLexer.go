package template

import (
	"log"
	"os"
	"regexp"

	. "LexGo/src"
	. "LexGo/src/bin"
	"LexGo/src/config"
	. "LexGo/src/regex"
	. "LexGo/src/tokens"
)

const Name = "LexGo"

func LexCodeFiles(filename string) (outputFilename string) {
	regex := CompileRegex()

	code, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	tokens := LexTokens(regex, &code)
	tokenset := NewTokenSet(tokens, filename)

	outputFilename = config.OUTPUT_FILENAME
	Write([]TokenSet{*tokenset}, outputFilename)
	return outputFilename
}

func CompileRegex() *Regex {
	regexFile, err := os.ReadFile(config.RULESET_REGEX)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(string(regexFile))
	return NewRegex(re)
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
	var tokenIDs = make([]string, 0)
	var values = make([]string, 0)
	tokens := make([]Token, 0)

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
	return &tokens
}
