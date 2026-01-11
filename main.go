package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func main() {
	csvFileName := flag.String("csv", "quiz.csv", "a csv file in the format of 'problem,solution'")
	limit := flag.Int("limit", 10, "Enter the time limit for the quiz in seconds. Defaults to 10")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file %v ", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse CSV File")
	}

	problems := parseLines(lines)
	// fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0

	answer := make(chan string)
	for i, p := range problems {
		fmt.Printf("Problem #%d %s ", i+1, p.q)

		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answer <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTest Ended! You've got %d out of %d", correct, len(problems))
			return
		case ans := <-answer:
			if ans == p.a {
				correct++
			}
		}
	}
	fmt.Printf("\nTest Ended! You've got %d out of %d", correct, len(problems))
}
