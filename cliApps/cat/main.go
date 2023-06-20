package main
import (
	"flag"
	"fmt"
	"bufio"
	"os"
	"log"
)
func main() {
	fileName := parseArgs()
	readFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan(){
		fmt.Println(fileScanner.Text())
	}
}

func parseArgs() (*string){
	filename:= flag.String("file", "main.go", "Specify the file name")
	flag.Parse()
	return filename
}
