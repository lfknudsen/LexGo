package bin

import (
	"fmt"
	"log"
	"os"
)

// AcceptTokens reads back the binary file-format the rest of the application saves its
// output to. It thus exists more for the purposes of testing and example than necessarily utility.
func AcceptTokens(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic(err)
		}
	}(file)
	bin := DecompileBinFile(file)
	fmt.Printf("Decoded the binary file:\n")
	bin.Print()
}
