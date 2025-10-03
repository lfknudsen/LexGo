package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"LexGo/src/bin"
	"LexGo/src/config"
	"LexGo/src/rules"
	"LexGo/template"
)

func main() {
	// Simple example of compiling rules, lexing code file, encoding tokens to binary,
	// and finally decoding and printing the binary file's contents.
	if len(os.Args) == 1 {
		_ = rules.CompileRulesetRegex(config.RULESET)
		tokenFile := template.LexCodeFiles("code.txt")
		bin.AcceptTokens(tokenFile)
		return
	}
	args := os.Args[1:]
	files := HandleCommandLineArguments(args)
	_ = rules.CompileRulesetRegex(config.RULESET)
	_ = template.LexCodeFiles(files...)
	return
}

func HandleCommandLineArguments(args []string) []string {
	var files []string
	if len(args) == 2 && args[0] == "decode" {
		bin.AcceptTokens(args[1])
		os.Exit(0)
	}
	for i := 0; i < len(args); i++ {
		if args[i][0] == '-' {
			if i < len(args)-1 {
				usedSubsequent, err := HandleCommandOption(args[i], args[i+1])
				if err != nil {
					os.Exit(1)
				}
				if usedSubsequent {
					i++
				}
			} else {
				err := HandleCommandOptionNoSubsequent(args[i])
				if err != nil {
					os.Exit(1)
				}
			}
		} else {
			files = append(files, args[i])
		}
	}
	return files
}

func HandleCommandOption(argument, subsequent string) (usedSubsequent bool, err error) {
	if strings.HasPrefix(argument, "--") {
		return HandleSingleOption(argument[2:], subsequent)
	}
	if strings.HasPrefix(argument, "-") {
		if len(argument) == 2 {
			return HandleSingleOption(argument[1:], subsequent)
		} else if len(argument) > 2 {
			for _, c := range argument[1:] {
				err := HandleConcatenatedOption(c)
				if err != nil {
					return false, err
				}
			}
			return false, nil
		}
	}
	return false, PrintHelp()
}

func HandleConcatenatedOption(argument rune) error {
	switch argument {
	case 'n':
		config.USE_BOM = false
	case 'u':
		config.USE_BOM = true
	case 'l':
		config.BYTE_ORDER = binary.LittleEndian
	case 'b':
		config.BYTE_ORDER = binary.BigEndian
	case 'p':
		config.OUTPUT_FORMAT = config.PLAINTEXT
	default:
		return PrintHelp()
	}
	return nil
}

func HandleSingleOption(argument, subsequent string) (usedSubsequent bool, err error) {
	switch strings.ToLower(argument) {
	case "no-bom", "n":
		config.USE_BOM = false
		return false, nil
	case "use-bom", "u": // default is true
		config.USE_BOM = true
		return false, nil
	case "endian", "e":
		switch subsequent {
		case "little", "little-endian":
			config.BYTE_ORDER = binary.LittleEndian
		case "big", "big-endian": // default is big endian
			config.BYTE_ORDER = binary.BigEndian
		case "native", "native-endian", "machine":
			config.BYTE_ORDER = binary.NativeEndian
		default:
			return true, PrintHelp()
		}
	case "rule", "rules", "ruleset", "r":
		config.RULESET = subsequent
	case "format", "f":
		switch subsequent {
		case "bin", "binary", "b": // default is binary
			config.OUTPUT_FORMAT = config.BINARY
		case "plain", "plaintext", "text", "p":
			config.OUTPUT_FORMAT = config.PLAINTEXT
		default:
			return true, PrintHelp()
		}
	case "output", "o":
		config.OUTPUT_FILENAME = subsequent
	}
	return true, nil
}

func HandleCommandOptionNoSubsequent(argument string) error {
	switch strings.ToLower(argument) {
	case "no-bom", "n":
		config.USE_BOM = false
	case "use-bom", "u": // default is true
		config.USE_BOM = true
	default:
		return PrintHelp()
	}
	return nil
}

func PrintHelp() error {
	fmt.Println(`
$ ./LexGo <options> <filename(s)>                       Tokenise code file(s) based on a ruleset
$ ./LexGo decode <filename>                             Pretty print contents of a binary file written by this programme. 

Options:
    --no-bom  -n                                        Do not use a BOM.
    --use-bom -u                                        Use a BOM (default)
    
    --endian -e     <endianness>                        Set the endianness.
                    little/little-endian
                    big/big-endian                      (default)
                    native/native-endian/machine
                    
    --rule -r       <ruleset filename>                  Specify ruleset filename (default is "ruleset.txt")
    
    --format -f     <format>                            Choose output format
                    bin/binary/b
                    plain/plaintext/text/p
    
    --output -o     <output filename>                   Set output filename.`)
	return nil
}
