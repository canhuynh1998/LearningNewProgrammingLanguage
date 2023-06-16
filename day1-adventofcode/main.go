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
	currentMax := 0.0
	resultMax := -math.MaxFloat64
	topThree := make([]int, 3)

    readFile, err := os.Open("puzzle_input.txt")

    if err != nil {
        log.Fatalln(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)


    for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == ""{
			resultMax = math.Max(resultMax, currentMax)
			compareValues(topThree, int(currentMax))
			currentMax = 0
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
	fmt.Println(topThree[0] + topThree[1] + topThree[2])

}

func compareValues( topThree []int, currentMax int) {
	if currentMax > topThree[0]{
		topThree[2] = topThree[1]
		topThree[1] = topThree[0]
		topThree[0] = currentMax
	}else if currentMax > topThree[1]{
		topThree[2] = topThree[1]
		topThree[1] = currentMax
	}else if currentMax > topThree[2]{
		topThree[2] = currentMax
	}
}