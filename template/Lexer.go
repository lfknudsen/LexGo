package template

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"LexGo/src"
)

const Name = "LexGo"

func OpenCodeFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	regexFile, err := os.ReadFile("in_re.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("About to compile...\n")
	re := regexp.MustCompile(string(regexFile))
	bufferedReader := bufio.NewReader(file)
	fileCopy, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	bReader := bufio.NewReader(fileCopy)
	runeReader := io.RuneReader(bufferedReader)
	regex := src.NewRegex(re)
	Lex(regex, runeReader, bReader)
}

func Lex(regex *src.Regex, reader io.RuneReader, bReader *bufio.Reader) src.TokenSet {
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

	fmt.Println("Beginning matching.")
	match := regex.FindReaderSubmatchIndex(reader)
	for i := 1; match != nil; i++ {
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
				fmt.Printf("%d: %s. L:%d; R:%d\n", i, name, left, right)
				var runeSz int
				var word []rune
				for readSoFar := 0; readSoFar < right-left; readSoFar += runeSz {
					var err error
					var runeRead rune
					runeRead, runeSz, err = bReader.ReadRune()
					if err == io.EOF {
						break
					} else if err != nil {
						log.Panic(err)
					}
					word = append(word, runeRead)
				}
				fmt.Printf("Word: %s.\n", string(word))
				break
				//fmt.Printf("Read %d bytes: %s\n", runeSz, string(word))
				//n, err := io.ReadFull(bReader, buffer)
				//if err != nil {
				//	return src.TokenSet{}
				//}
				//tokens = append(tokens, )
			}
		}
		match = regex.FindReaderSubmatchIndex(reader)
	}
	return src.TokenSet{}
}
