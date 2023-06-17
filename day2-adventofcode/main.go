package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
func main() {
	// readFile, err := os.Open("puzzle_input.txt")
	readFile, err := os.Open("test_input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)

	scores := constructScoreMap()

	currentScore := 0
	for filescanner.Scan(){
		currentScore += scores[filescanner.Text()]
	}
	fmt.Println(currentScore)

}

func constructScoreMap() (map[string] int) {
	return map[string]int{
		"A X": 4, 
		"A Y": 8, 
		"A Z": 3, 
		"B X": 1,
		"B Y": 5,
		"B Z": 9, 
		"C X": 7, 
		"C Y": 2, 
		"C Z": 6,
	}

}