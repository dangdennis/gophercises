// flag, csv, time, os, strings
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type problem struct {
	Operation string
	Solution  int
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
	flag.IntVar(&flagTime, "time", 10, "test duration")
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

		solution, _ := strconv.Atoi(line[1])
		quiz = append(quiz, problem{
			Operation: line[0],
			Solution:  solution})
	}

	var correct int
	for _, problem := range quiz {
		fmt.Print(problem.Operation, "=")

		var answer int
		fmt.Scanf("%d\n", &answer)

		if answer == problem.Solution {
			fmt.Print("Correct\n")
			correct++
		} else {
			fmt.Print("Incorrect\n")
		}
	}

	fmt.Println("You got", correct, "correct out of", len(quiz))
}
