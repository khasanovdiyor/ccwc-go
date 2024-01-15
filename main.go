package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	cFlag := false
	flag.CommandLine.BoolVar(&cFlag, "c", false, "outputs the char count in a text file")
	lFlag := false
	flag.CommandLine.BoolVar(&lFlag, "l", false, "outputs the char count in a text file")
	flag.Parse()

	fileName := flag.Arg(0)
	if len(fileName) == 0 {
		fmt.Println("Please provide a file name as an argument, example: text.txt")
		os.Exit(1)
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Error getting file information:", err)
		return
	}

	if cFlag {
		fmt.Println(fileInfo.Size(), fileName)
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error when opening the file", err)
		return
	}

	if lFlag {
		lineCount := len(strings.Split(string(file), "\n"))
		fmt.Println("\t", lineCount, fileName)
	}
}
