package main

import (
	"bufio"
	// "unicode/utf8"
	"fmt"
	"log"
	"os"
	"strings"
)


func main(){
	cmdArgs, err := readCmdArgs()

	if err != nil  || len(cmdArgs) < 2 {
		log.Fatal(err)
	}
	result, _ := count(cmdArgs[1], cmdArgs[2])

	fmt.Println(result)

}

func count( flag string, filename string) (string, error) {
	fileContent := readLinesInFile(filename)
	defer fileContent.Close()
	if flag == "-l"{
		return fmt.Sprintf("\t%d %s", countAllLines(fileContent), filename), nil
	}else if flag == "-c" {
		return fmt.Sprintf("\t%d %s", countAllBytes(fileContent), filename), nil
	}else if flag == "-w" {
		return fmt.Sprintf("\t%d %s", countAllWords(fileContent), filename), nil
	}else if flag == "-m"{
		return fmt.Sprintf("\t%d %s", countAllChars(fileContent), filename), nil
	}
	return fmt.Sprintf("\t%d %d %d %s", countAllLines(fileContent), countAllWords(fileContent), countAllWords(fileContent), filename), nil
}

func countAllLines(fileContent *os.File) int {
	scanner := bufio.NewScanner(fileContent)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan(){
		count ++
	}
	return count
}
func countAllBytes(fileContent *os.File) int {
	scanner := bufio.NewScanner(fileContent)
	scanner.Split(bufio.ScanBytes)
	count := 0
	for scanner.Scan(){
		count ++
	}
	return count
}

func countAllWords(fileContent *os.File) int {
	scanner := bufio.NewScanner(fileContent)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan(){
		count ++
	}
	return count
}

func countAllChars(fileContent *os.File) int {
	scanner := bufio.NewScanner(fileContent)
	scanner.Split(bufio.ScanRunes)
	count := 0
	for scanner.Scan(){
		count ++
	}
	return count

}

func readCmdArgs() ( []string, error){
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