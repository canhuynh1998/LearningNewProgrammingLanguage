package main
import (
	"os"
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
)
func main() {

    readFile, err := os.Open("puzzle_input.txt")

    if err != nil {
        log.Fatalln(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
	currentMax := 0.0
	resultMax := -math.MaxFloat64
	topThree := make([]int, 3)
    for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == ""{
			resultMax = math.Max(resultMax, currentMax)
			currentMax = 0.0
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		
		if err != nil{
			log.Fatalln(err)
		}
		currentMax += num
    }
  
    readFile.Close()
	fmt.Println(resultMax)

}