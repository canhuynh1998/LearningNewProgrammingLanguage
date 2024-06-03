package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
)

type Args struct {
	Content string
	Flags   map[string]bool
}

func main() {
	flags := ParseFlag()
	content := ReadFileContent()
var wg sync.WaitGroup
	args := Args{content, flags}
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	for i := 0; i < 2; i++{
		wg.Add(1)
        go func() {
            defer wg.Done()
            Execute(args, client, err, i)
        }()
	}
	  wg.Wait()
}

func Execute(args Args, client *rpc.Client, err error, idx int){

	var reply string
	fmt.Println(idx)
	err = client.Call("Wc.Count", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Report: %s\n", reply)
}

func ReadFileContent() string {
	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	buffer, _ := io.ReadAll(file)
	return string(buffer)
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
