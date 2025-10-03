package main

import (
	"fmt"
	"log"
	"os"

	"LexGo/src"
	"LexGo/src/bin"
	"LexGo/template"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		filename := src.CompileRulesetRegex("ruleset.txt")
		f, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of file with compiled regexp:\n%s\n", f)
		tokenFile := template.OpenCodeFile("code.txt")
		bin.AcceptTokens(tokenFile)
		return
	}

}
