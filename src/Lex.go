package src

import (
	"fmt"
	"go/scanner"
	"go/token"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"LexGo/src/regex"
)

func Lex(filename string) (resultingFilename string) {
	ruleset, err := ReadSpecs(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ruleset.String())
	compiled := ruleset.Compile()
	fmt.Println("Compiled regexp:")
	fmt.Println(compiled.String())

	basename := path.Base(filename)
	ext := path.Ext(basename)
	name, _ := strings.CutSuffix(basename, ext)
	outputFilename := path.Join(path.Dir(filename), name+"_re"+ext)
	outputFile, err := os.Create(outputFilename)
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(outputFile)

	n, err := outputFile.WriteString(compiled.String())
	if err != nil {
		log.Panic(err)
	}
	if n != len(compiled.String()) {
		log.Panicf("%d bytes written to %s", n, outputFilename)
	}
	log.Printf("%d bytes written to %s", n, outputFilename)

	return outputFilename
}

func ErrorHandler(pos token.Position, msg string) {
	log.Panicf("Lex: Error at %s: %s", pos, msg)
}

type ScannedToken struct {
	pos token.Position
	tok token.Token
	lit string
}

func (t ScannedToken) String() string {
	return fmt.Sprintf("%d, %d: %s - %s", t.pos.Line, t.pos.Column, t.tok, t.lit)
}

func SplitTokens(r rune) (shouldBeSplit bool) {
	if r == '\n' {
		return true
	}
	return false
}

// ReadSpecs lexes an input file with the syntax <Name><whitespace><Regexp><\n>(repeat)
// and produces the resulting Ruleset.
func ReadSpecs(filename string) (*Ruleset, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	ruleset := new(Ruleset)

	// ruleString := `(?m)^\s*(?<COMMENT_BLOCK>/\*(?:[^*][^/])*\*/)|(?<COMMENT_LINE>#[^\n])|(?:(?P<ID>\S+)[\s&^\n]+(?P<REGEX>(?:[^\n]*\S))[\s&^\n]*)$`
	// ruleRegexp := regexp.MustCompile(ruleString)
	rex := regex.LoadRegex("Expressions/ReadRulesheet.txt")
	match := rex.FindSubmatchIndex(contents)
	for i := 1; match != nil; i++ {
		leftID, rightID := rex.Group(regex.ID, match)
		// log.Printf("ID indices %d, %d.\n", leftID, rightID)
		leftRegexp, rightRegexp := rex.Group(regex.REGEX, match)
		leftEncoding, rightEncoding := rex.Group(regex.ENCODING, match) // Optional
		// log.Printf("Regex indices %d, %d.\n", leftRegexp, rightRegexp)
		if leftID == rightID || leftRegexp == rightRegexp {
			left, right := rex.Group(regex.COMMENT_BLOCK, match)
			// log.Printf("Comment block indices %d, %d.\n", left, right)
			if left != -1 && left != right {
				contents = contents[right:]
			} else {
				left, right = rex.Group(regex.COMMENT_LINE, match)
				// log.Printf("Comment line indices %d, %d.\n", left, right)
				if left != -1 && left != right {
					// log.Printf("Content length: %d\n", len(contents))
					contents = contents[right:]
				} else {
					left, right = rex.Group(regex.MISTAKE, match)
					// log.Printf("Mistake indices %d, %d.\n", left, right)
					if left != -1 && left != right {
						// log.Printf("Error on row %d, columns %d-%d. The term '%v' could not be recognised.",
						//	i, left, right, string(contents[left:right]))
						contents = contents[right:]
					} else {
						log.Printf("%v\n", contents[:50])
						log.Panicf("Lex: Error on line %d", i)
					}
				}
			}
		} else {
			log.Printf("Successfully found a pair on row %d\n", i)
			id := string(contents[leftID:rightID])
			en := string(contents[leftEncoding:rightEncoding])
			re := string(contents[leftRegexp : rightRegexp-1]) // Trimming away newline
			_ = ruleset.Add(NewEncodedRule(id, re, en))
			fmt.Printf("ID: %v, Regex: %v\n", id, re)
			contents = contents[rightRegexp:]
		}
		match = rex.FindSubmatchIndex(contents)
	}
	fmt.Println(ruleset.String())
	return ruleset, nil
}

func UseScanner(filename string) {
	code, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	tokenScanner := scanner.Scanner{}
	fileset := token.NewFileSet()
	file := fileset.AddFile(filename, fileset.Base(), len(code))
	tokenScanner.Init(file, code, ErrorHandler, 0)
	scanned := make([]ScannedToken, 500)
	var pos token.Pos
	var tok token.Token
	var literal string

	for tokenScanner.ErrorCount == 0 && tok != token.EOF {
		pos, tok, literal = tokenScanner.Scan()
		scanned = append(scanned, ScannedToken{file.Position(pos), tok, literal})
		if tok == token.EOF {
			break
		}
	}

	fmt.Println("Scanned " + strconv.Itoa(len(scanned)) + " tokens")
	for _, scannedToken := range scanned {
		fmt.Println(scannedToken.String())
	}
}
