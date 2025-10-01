package template

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"regexp"

	"LexGo/src"
	"LexGo/src/bin"
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
	tokens := Lex(regex, &code)
	tokenset := bin.NewTokenSet(tokens, filename)
	bin.Write(tokenset, filename+"_out.txt")
}

func Lex(regex *src.Regex, code *[]byte) *[]src.Token {
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
	var values []string = make([]string, 0)
	tokens := make([]src.Token, 0)

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
				token := src.Token{
					ID:    byte(i),
					Value: make([]byte, 0),
				}
				if rule.Encoding == reflect.Int {
					n, err := binary.Encode(token.Value,
						src.BYTE_ORDER,
						BytesToInt(value))
					if err != nil {
						log.Fatal(err)
					}
					token.ValueLength = uint16(n)
				} else {
					token.Value = value
					token.ValueLength = uint16(len(value))
				}
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

func BytesToInt(buffer []byte) int64 {
	var number int64 = 0
	for i := 0; i < len(buffer); i++ {
		number *= 10
		number += int64(buffer[i])
	}
	return number
}

func FitsInto(value uint64) reflect.Kind {
	if value <= math.MaxInt8 {
		return reflect.Int8
	}
	if value <= math.MaxInt16 {
		return reflect.Int16
	}
	if value <= math.MaxInt32 {
		return reflect.Int32
	}
	if value <= math.MaxInt {
		return reflect.Int
	}
	if value <= math.MaxInt64 {
		return reflect.Int64
	}
	return math.MaxUint64
}
