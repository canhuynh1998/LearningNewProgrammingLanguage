- To read in txt file, use `bufio`
```go
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
    // bufio.ScanLines is a function that splits lines from the file

    for fileScanner.Scan(){
        // ...
    }
```