package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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
	_ = csvFileName
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

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d %s ", i+1, p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			correct++
		}
	}
	fmt.Printf("You've got %d out of %d", correct, len(problems))
}
