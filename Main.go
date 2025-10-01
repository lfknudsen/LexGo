package main

import (
	"fmt"
	"log"
	"os"

	"LexGo/src"
	"LexGo/template"
)

func main() {
	// var err error
	/*
		t := NewToken(8, 7, []byte("Hello, my old friend!"), "Test filename", 3, 49)
		fmt.Println(t.String())
		f, err := os.Create("test.txt")
		if err != nil {
			log.Fatal(err)
		}
		t.Write(f)
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
		open, err := os.Open("test.txt")
		if err != nil {
			log.Fatal(err)
		}
		b := make([]byte, 100)
		n, err := open.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Read " + strconv.Itoa(n) + " bytes from " + open.Name())
		err = open.Close()
		if err != nil {
			return
		}
		returned, err := ReadToken(b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(returned.Equals(t))
	*/
	filename := src.Lex("in.txt")
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", f)
	template.OpenCodeFile("code.txt")
	// regexp := regexp.MustCompile(string(f))

}
