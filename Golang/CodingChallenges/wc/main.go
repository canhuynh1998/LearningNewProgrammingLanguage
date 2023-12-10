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
		if len(cmdArgs) == 2{
			result = countFromFile("", cmdArgs[1])
		}else{
			result = countFromFile(cmdArgs[1], cmdArgs[2])
		}
	}else {
		fmt.Println(cmdArgs)
	}

	

	fmt.Println(result)

}

func countFromFile(flag string, filename string) (string) {
	fileContent := readLinesInFile(filename)
	
	
	if flag == "-l" {
		return fmt.Sprintf("\t%d %s", countAllLines(*fileContent), filename)
	} else if flag == "-c" {
		return fmt.Sprintf("\t%d %s", countAllBytes(*fileContent), filename)
	} else if flag == "-w" {
		return fmt.Sprintf("\t%d %s", countAllWords(*fileContent), filename)
	} else if flag == "-m" {
		return fmt.Sprintf("\t%d %s", countAllChars(*fileContent), filename)
	}
	return fmt.Sprintf("\t%d %d %d %s", countAllLines(*fileContent), countAllWords(*fileContent),countAllWords(*fileContent), filename)
}

func countHelper(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		count++
	}
	
	return count
}

func countAllLines(fileContent os.File) int {
	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanLines)
	return countHelper(scanner)
}

func countAllBytes(fileContent os.File) int {

	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanLines)
	return countHelper(scanner)
}

func countAllWords(fileContent os.File) int {
	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanLines)
	return countHelper(scanner)
}

func countAllChars(fileContent os.File) int {
	scanner := bufio.NewScanner(&fileContent)
	scanner.Split(bufio.ScanLines)
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

func readLinesInFile(filename string) *os.File {
	fileContent, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fileContent
}
