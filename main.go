package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	printByteCount := false
	flag.CommandLine.BoolVar(&printByteCount, "c", false, "outputs the byte count in a text file")
	printLineCount := false
	flag.CommandLine.BoolVar(&printLineCount, "l", false, "outputs the line count in a text file")
	printWordCount := false
	flag.CommandLine.BoolVar(&printWordCount, "w", false, "outputs the word count in a text file")
	printCharCount := false
	flag.CommandLine.BoolVar(&printCharCount, "m", false, "outputs the char count in a text file")
	flag.Parse()

	var fileContents []byte

	fileName := flag.Arg(0)
	if len(fileName) == 0 {
		var err error
		fileContents, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		var err error
		fileContents, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error when opening the file", err)
			return
		}
	}

	if !printByteCount && !printLineCount && !printWordCount && !printCharCount {
		printByteCount = true
		printLineCount = true
		printWordCount = true
	}

	fileStats := getStats(fileContents)
	var arr []string

	if printByteCount {
		arr = append(arr, fmt.Sprint(fileStats.byteCount))
	}

	if printLineCount {
		arr = append(arr, fmt.Sprint(fileStats.lineCount))
	}

	if printWordCount {
		arr = append(arr, fmt.Sprint(fileStats.wordCount))
	}

	if printCharCount {
		arr = append(arr, fmt.Sprint(fileStats.charCount))
	}

	if len(fileName) > 0 {
		arr = append(arr, fileName)
	}

	fmt.Println(getResultStr(arr))
}

func getBytesCount(fileContents []byte) uint64 {
	return uint64(len(fileContents))
}

func getLinesCount(fileContents []byte) uint64 {
	return uint64(len(strings.Split(string(fileContents), "\n")))
}

func getWordsCount(fileContents []byte) uint64 {
	return uint64(len(strings.Fields(string(fileContents))))
}

func getCharsCount(fileContents []byte) uint64 {
	return uint64(len([]rune(string(fileContents))))
}

func getStats(fileContents []byte) stats {
	return stats{byteCount: getBytesCount(fileContents), lineCount: getLinesCount(fileContents), wordCount: getWordsCount(fileContents), charCount: getCharsCount(fileContents)}
}

func getResultStr(arr []string) string {
	return strings.Join(arr, "\t")
}

type stats struct {
	byteCount uint64
	lineCount uint64
	wordCount uint64
	charCount uint64
}
