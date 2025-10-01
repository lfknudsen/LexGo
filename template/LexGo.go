package template

import (
	"fmt"
	"os"
)

func main() {
	os.Args = os.Args[1:]
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run LexGo.go <filename>")
		os.Exit(1)
	}
	OpenCodeFile(os.Args[1])
}
