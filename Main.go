package main

import (
	"os"

	"LexGo/src"
	"LexGo/src/bin"
	"LexGo/template"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		_ = src.CompileRulesetRegex("ruleset.txt")
		tokenFile := template.OpenCodeFile("code.txt")
		bin.AcceptTokens(tokenFile)
		return
	}

}
