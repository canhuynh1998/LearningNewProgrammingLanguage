package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cmdArgs, err := readCmdArgs()
	if err != nil || len(cmdArgs) < 2 {
		log.Fatal(err.Error())
	}
	result := ""
	if cmdArgs[0] == "ccwc" {
		if len(cmdArgs) == 2 {
			result = countFromFile("", cmdArgs[1])
		} else {
			result = countFromFile(cmdArgs[1], cmdArgs[2])
		}
	} else {
		fmt.Println(cmdArgs)
	}

	fmt.Println(result)

}

func countFromFile(flag string, filename string) string {
	if flag == "-l" {
		return fmt.Sprintf("\t%d %s", countAllLines(filename), filename)
	} else if flag == "-c" {
		return fmt.Sprintf("\t%d %s", countAllBytes(filename), filename)
	} else if flag == "-w" {
		return fmt.Sprintf("\t%d %s", countAllWords(filename), filename)
	} else if flag == "-m" {
		return fmt.Sprintf("\t%d %s", countAllChar(filename), filename)
	} 
	return fmt.Sprintf("\t%d %d %d %s", countAllLines(filename), countAllWords(filename), countAllBytes(filename), filename)
}

func countHelper(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func countAllLines(filename string) int {
	fileContent := readLinesInFile(filename)
	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanLines)
	return countHelper(scanner)
}

func countAllBytes(filename string) int {
	fileContent := readLinesInFile(filename)
	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanBytes)
	return countHelper(scanner)
}

func countAllWords(filename string) int {
	fileContent := readLinesInFile(filename)
	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanWords)
	return countHelper(scanner)
}

func countAllChar(filename string) int {
	fileContent := readLinesInFile(filename)
	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanRunes)
	return countHelper(scanner)
}

func readCmdArgs() ([]string, error) {
	fmt.Print(">")
	scanner := bufio.NewReader(os.Stdin)
	cmdInput, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	cmdInput = strings.TrimSuffix(cmdInput, "\n")
	return strings.Split(cmdInput, " "), nil
}

func readLinesInFile(filename string) os.File {
	fileContent, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return *fileContent
}
