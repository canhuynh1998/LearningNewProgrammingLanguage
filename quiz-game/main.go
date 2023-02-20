package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
    question string
    solution string
}

func main() {
    filename, timer := parseArg()
    csvFile, err := os.Open(filename)

    if err != nil {
        log.Fatalln(err)
    }
    csvReader := csv.NewReader(csvFile)
    data, err := csvReader.ReadAll()

    questions := parseQuestion(data)
    counts := answerQuestion(questions, timer)


    fmt.Printf("\nTotal score: %d\n", counts)

}

func answerQuestion(questions []Question, timer int) (int) {
    counts := 0
    timeLimit := time.NewTimer(time.Duration(timer)*time.Second)
    for idx , q := range questions {
        fmt.Printf("Problem %d: %s = ", idx + 1, q.question)
        answerChannel := make(chan string)
        go func() {
            var answer string
            fmt.Scanf("%s\n", &answer)
            answer = strings.TrimSpace(answer)
            answerChannel <- answer
        } ()
        select {
        case <- timeLimit.C:
            return counts
        case answer := <- answerChannel:
            counts += compareAnswer(q.solution, answer)
        }
    }

    return counts
}


func compareAnswer(solution string, answer string) (int) {
    if answer == solution {
        return 1
    }
    return 0
}
func parseQuestion(data [][] string) []Question {
    questions := make([]Question, len(data))

    for idx, question := range data {
        current := Question{question[0], question[1]}
        questions[idx] = current
    }

    return questions
}

func parseArg() (string, int) {
    fileArg := flag.String("file", "problem.csv", "Specify the file name")
    time := flag.Int("time", 30, "Time limit for 1 question")
    flag.Parse()
    return *fileArg, *time
}