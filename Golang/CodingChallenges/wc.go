package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Result struct {
	name  string
	bytes int
	lines int
	words int
	chars int
}

func main() {
	flags := ParseFlag()
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fi.Size() > 0 {
		Execution("")
	} else {
		var computeResult []Result
		for _, filename := range flag.Args() {
			currentResult := Execution(filename)
			computeResult = append(computeResult, currentResult)
			Display(currentResult, flags)
		}
		Display(computeTotal(computeResult), flags)

	}

}
func ParseFlag() map[string]bool {
	countingBytes := flag.Bool("c", false, "count bytes from the input")
	countingLines := flag.Bool("l", false, "count lines from the input")
	countingWords := flag.Bool("w", false, "count words from the input")
	countingChars := flag.Bool("m", false, "count character from the input")

	flag.Parse()
	flags := make(map[string]bool)
	flags["countBytes"] = *countingBytes
	flags["countLines"] = *countingLines
	flags["countWords"] = *countingWords
	flags["countChars"] = *countingChars
	return flags
}

func BuildContent(filename string) string {
	if filename == "" {
		return CreateInput(os.Stdin)
	}
	file, err := os.Open("./" + filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	return CreateInput(file)
}

func CreateInput(file io.Reader) string {
	buffer, _ := io.ReadAll(file)
	return string(buffer)
}

func Execution(filename string) Result {
	content := BuildContent(filename)
	bChan := make(chan int)
	lChan := make(chan int)
	wChan := make(chan int)
	cChan := make(chan int)

	go CountLines(content, lChan)
	go CountWords(content, wChan)
	go CountBytes(content, bChan)
	go CountChars(content, cChan)
	return Result{name: filename, lines: <-lChan, words: <-wChan, bytes: <-bChan, chars: <-cChan}
}


func Display(result Result, flags map[string] bool){
	if flags["countBytes"]{
		fmt.Printf("  %d %s\n", result.bytes, result.name)
	} else if flags["countLines"] {
		fmt.Printf("  %d %s\n", result.lines, result.name)
	} else if flags["countWords"] {
		fmt.Printf("  %d %s\n", result.words, result.name)
	} else if flags["countChars"] {
		fmt.Printf("  %d %s\n", result.chars, result.name)
	}else{
		fmt.Printf("    %d   %d   %d %s\n", result.lines, result.words, result.bytes, result.name)
	}
}

func computeTotal(results []Result) Result{
	totalBytes := 0
	totalLines := 0
	totalWords := 0
	totalChars := 0
	for _, result := range(results){
		totalBytes += result.bytes
		totalLines += result.lines
		totalWords += result.words
		totalChars += result.chars
	}
	return Result{name:"total", bytes: totalBytes, lines: totalLines, words: totalWords, chars: totalChars}
}

func CountBytes(content string, bChan chan int) {
	bChan <- CountHepler(strings.NewReader(content), bufio.ScanBytes)
}

func CountLines(content string, lChan chan int) {
	lChan <- CountHepler(strings.NewReader(content), bufio.ScanLines)
}

func CountWords(content string, wChan chan int) {
	wChan <- CountHepler(strings.NewReader(content), bufio.ScanWords)
}

func CountChars(content string, cChan chan int) {
	cChan <- CountHepler(strings.NewReader(content), bufio.ScanRunes)
}

func CountHepler(file io.Reader, spiltFunc bufio.SplitFunc) int {
	count := 0
	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		return -1
	}
	scanner.Split(spiltFunc)
	for scanner.Scan() {
		count += 1
	}
	return count
}
