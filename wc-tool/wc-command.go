package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	check(err)

	return data
}

func output(command string, filePath string) {
	switch cases := command; cases {
	case "-c":
		data := ReadFile(filePath)
		fmt.Println(len(data), filePath)

	case "-l":
		count := 0
		data := ReadFile(filePath)
		lineSep := []byte{'\n'}
		count += bytes.Count(data, lineSep)
		fmt.Println(count, filePath)

	case "-w":
		data := ReadFile(filePath)
		wordsWithoutNewLine := strings.Replace(string(data), "\n", "", -1)
		words := strings.Fields(wordsWithoutNewLine)
		fmt.Println(len(words), filePath)

	case "-m":
		data := ReadFile(filePath)
		fmt.Println(utf8.RuneCountInString(string(data)), filePath)

	default:
		count := 0
		data := ReadFile(filePath)
		lineSep := []byte{'\n'}
		count += bytes.Count(data, lineSep)
		wordsWithoutNewLine := strings.Replace(string(data), "\n", "", -1)
		words := strings.Fields(wordsWithoutNewLine)
		fmt.Println(len(data), count, len(words), filePath)
	}
}

func main() {
	if len(os.Args) == 4 {
		baseCommand := os.Args[1]
		command := os.Args[2]
		filePath := os.Args[3]
		if baseCommand == "ccwc" {
			output(command, filePath)
		} else {
			check(errors.New("command incorrect"))
		}
	}
	if len(os.Args) == 3 {
		baseCommand := os.Args[1]
		filePath := os.Args[2]
		if baseCommand == "ccwc" {
			output("-default", filePath)
		} else {
			check(errors.New("command incorrect"))
		}
	}
}
