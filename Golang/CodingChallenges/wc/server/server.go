package main

import (
	// "errors"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

type Args struct {
	Content string
	Flags map[string]bool
}

type Result struct {
	bytes int
	lines int
	words int
	chars int
}

type Wc string

func (t *Wc) Count(args *Args, reply *string) error {
	result := Execution_(args.Content)
	if args.Flags["countBytes"] {
		*reply = fmt.Sprintf("  %d bytes\n", result.bytes)
	} else if args.Flags["countLines"] {
		*reply = fmt.Sprintf("  %d lines\n", result.lines)
	} else if args.Flags["countWords"] {
		*reply = fmt.Sprintf("  %d words\n", result.words)
	} else if args.Flags["countChars"] {
		*reply = fmt.Sprintf("  %d chars\n", result.chars)
	} else {
		*reply = fmt.Sprintf("    %-5d lines   %-5d words   %-5d bytes\n", result.lines, result.words, result.bytes)
	}
	return nil
}

func main() {
	// Create a new RPC server
	wc := new(Wc)
	rpc.Register(wc)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)

}

func Execution_(content string) Result {
	return Result{lines: CountLines_(content), words:  CountWords_(content), bytes: CountBytes_(content), chars: CountChars_(content)}
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

func CountBytes_(content string) int {
	return CountHepler(strings.NewReader(content), bufio.ScanBytes)
}

func CountLines_(content string) int {
	return CountHepler(strings.NewReader(content), bufio.ScanLines)
}

func CountWords_(content string) int {
	return CountHepler(strings.NewReader(content), bufio.ScanWords)
}

func CountChars_(content string) int {
	return CountHepler(strings.NewReader(content), bufio.ScanRunes)
}