package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	cmdArgs := parseArg()
	fromStdIn := false

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		fromStdIn = true
	} else {
		fromStdIn = false
	}

	fmt.Println(countFromFile(cmdArgs[0], cmdArgs[1], fromStdIn))
}

func parseArg() []string {
	args := make([]string, 2)
	cmd := os.Args[1:]
	if len(cmd) == 0 {
		return args
	}
	flag := createFlagMap()

	for i := 0; i < len(cmd); i++ {
		arg := cmd[i]
		_, ok := flag[arg]
		if ok {
			args[0] = arg
		} else {
			args[1] = arg
		}
	}
	return args
}

func createFlagMap() map[string]bool {
	flag := make(map[string]bool)
	flag["-l"] = true
	flag["-m"] = true
	flag["-w"] = true
	flag["-c"] = true
	return flag
}

func countFromFile(flag string, filename string, fromStdIn bool) string {
	displayName := ""
	if !fromStdIn {
		displayName = filename
	}

	return displayInformation(flag, filename, displayName, fromStdIn)

}

func displayInformation(flag string, filename string, displayName string, fromStdIn bool) string {
	if flag == "-l" {
		return fmt.Sprintf("\t%d %s", countAllLines(filename, fromStdIn), displayName)
		// return fmt.Sprintf("\t%d %s", countAllLines(filename), filename)
	} else if flag == "-c" {
		return fmt.Sprintf("\t%d %s", countAllBytes(filename, fromStdIn), displayName)
	} else if flag == "-w" {
		return fmt.Sprintf("\t%d %s", countAllWords(filename, fromStdIn), displayName)
	} else if flag == "-m" {
		return fmt.Sprintf("\t%d %s", countAllChar(filename, fromStdIn), displayName)
	}
	return fmt.Sprintf("\t%d %d %d %s", countAllLines(filename, fromStdIn), countAllWords(filename, fromStdIn), countAllBytes(filename, fromStdIn), displayName)
}

func countHelper(scanner *bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func readLinesInFileReader(filename string) io.Reader {
	fileContent, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fileContent
}

func countAllLines(filename string, fromStdIn bool) int {
	fmt.Println(filename)
	var fileContent io.Reader
	if fromStdIn {
		fileContent = os.Stdin
	} else {
		fileContent = readLinesInFileReader(filename)
	}
	scanner := bufio.NewScanner(fileContent)
	scanner.Split(bufio.ScanLines)
	return countHelper(scanner)
}

func countAllBytes(filename string, fromStdIn bool) int {
	var fileContent io.Reader
	if fromStdIn {
		fileContent = os.Stdin
	} else {
		fileContent = readLinesInFileReader(filename)
	}
	scanner := bufio.NewScanner(fileContent)
	scanner.Split(bufio.ScanBytes)
	return countHelper(scanner)
}

func countAllWords(filename string, fromStdIn bool) int {
	fmt.Println(filename)
	var fileContent io.Reader
	if fromStdIn {
		fileContent = os.Stdin
	} else {
		fileContent = readLinesInFileReader(filename)
	}
	scanner := bufio.NewScanner(fileContent)
	scanner.Split(bufio.ScanWords)
	return countHelper(scanner)
}

func countAllChar(filename string, fromStdIn bool) int {
		var fileContent io.Reader
	if fromStdIn {
		fileContent = os.Stdin
	} else {
		fileContent = readLinesInFileReader(filename)
	}
	scanner := bufio.NewScanner(fileContent)
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
