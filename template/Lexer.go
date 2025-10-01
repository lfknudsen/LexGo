package template

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"LexGo/src"
)

const Name = "LexGo"

func OpenCodeFile(filename string) {
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
	regex := src.NewRegex(re)
	Lex(regex, code)
}

func Lex(regex *src.Regex, code []byte) src.TokenSet {
	//tokens := make([]src.Token, 500)
	//buffer := make([]byte, 1024)

	ruleset := src.Decompile(regex.Src())
	if ruleset == nil {
		log.Panic("Ruleset is nil")
	}
	names := ruleset.Names()

	fmt.Println("Subexpression names:")
	for idx, name := range *names {
		fmt.Printf("%d: %s\n", idx, name)
	}

	var tokenIDs []string = make([]string, 0)
	var values [][]byte = make([][]byte, 0)

	fmt.Println("Beginning matching.")
	matches := regex.FindAllSubmatchIndex(code)
	for _, match := range matches {
		for _, name := range *names {
			if name == "" {
				continue
			}
			idx := regex.SubNames[name]
			if idx == -1 {
				continue
			}
			left, right := match[idx*2], match[idx*2+1]
			if left != -1 && left != right {
				value := code[left:right]
				values = append(values, value)
				tokenIDs = append(tokenIDs, name)
				break
			}
		}
	}
	fmt.Println("Result of pattern matching on code text:")
	for i, value := range values {
		fmt.Println(tokenIDs[i] + ":    " + string(value))
	}
	return src.TokenSet{}
}
