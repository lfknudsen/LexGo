package rules

import (
	"log"
	"os"
	"path"

	"LexGo/src/config"
	"LexGo/src/regex"
)

func CompileRulesetRegex(filename string) (resultingFilename string) {
	ruleset, err := ReadSpecs(filename)
	if err != nil {
		log.Fatal(err)
	}

	compiled := ruleset.Compile()

	outputFilename := path.Join(path.Dir(filename), config.RULESET_REGEX)
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
		log.Panicf("%d bytes written to %s, but expected %d",
			n, outputFilename, len(compiled.String()))
	}

	return outputFilename
}

// ReadSpecs lexes an input file with the syntax <Name><whitespace><Regexp><\n>(repeat)
// and produces the resulting Ruleset.
func ReadSpecs(filename string) (*Ruleset, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	ruleset := new(Ruleset)
	rex := regex.LoadRegex("Expressions/ReadRuleset.txt")
	match := rex.FindSubmatchIndex(contents)
	for i := 1; match != nil; i++ {
		leftID, rightID := rex.Group(regex.ID, match)
		leftRegexp, rightRegexp := rex.Group(regex.REGEX, match)
		if leftID == rightID || leftRegexp == rightRegexp {
			left, right := rex.Group(regex.COMMENT_BLOCK, match)
			if left != -1 && left != right {
				contents = contents[right:]
			} else {
				left, right = rex.Group(regex.COMMENT_LINE, match)
				if left != -1 && left != right {
					contents = contents[right:]
				} else {
					left, right = rex.Group(regex.MISTAKE, match)
					if left != -1 && left != right {
						contents = contents[right:]
					} else {
						log.Printf("%v\n", contents[:50])
						log.Panicf("CompileRulesetRegex: Error on line %d", i)
					}
				}
			}
		} else {
			id := string(contents[leftID:rightID])
			re := string(contents[leftRegexp : rightRegexp-1]) // Trimming away newline
			_ = ruleset.Add(NewRule(id, re))
			contents = contents[rightRegexp:]
		}
		match = rex.FindSubmatchIndex(contents)
	}
	return ruleset, nil
}
