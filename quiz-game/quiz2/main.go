package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type problem struct {
	Operation string
	Solution  string
}

var (
	flagFilePath string
	flagRandom   bool
	flagTime     int
	wg           sync.WaitGroup
)

func init() {
	flag.StringVar(&flagFilePath, "file", "problems.csv", "path/to/csv_file")
	flag.BoolVar(&flagRandom, "random", true, "randomize order of questions")
	flag.IntVar(&flagTime, "time", 1000, "test duration")
	flag.Parse()
}

func main() {
	fmt.Println("Quiz begins now:")

	csvFile, err := os.Open(flagFilePath)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	if err != nil {
		log.Fatal(err)
	}

	var quiz []problem

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		quiz = append(quiz, problem{
			Operation: line[0],
			Solution:  line[1]})
	}

	timer := time.NewTimer(time.Duration(flagTime) * time.Second)

	var correct int
	for _, problem := range quiz {
		fmt.Printf("%s= ", problem.Operation)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(quiz))
			return
		case answer := <-answerCh:
			if answer == problem.Solution {
				correct++
			}
		}
	}

	fmt.Printf("You got %d correct out of %d", correct, len(quiz))
}
