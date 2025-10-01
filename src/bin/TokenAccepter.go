package bin

import (
	"log"
	"os"
)

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

	_ = DecompileBinFile(file)
}
