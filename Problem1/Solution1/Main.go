package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("time", 3, "Time limit for the quiz")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s.csv.\n", *csvFilename))
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}
	problems := parseLines(lines)
	correct := 0
	fmt.Println("Press enter to Start timer")
	var pressKey string
	fmt.Scanf("%s", pressKey)
	go startTimer(*timeLimit)
	for i, p := range problems {
		if printProblem(p, i+1) {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func startTimer(seconds int) {
	timer1 := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer1.C
	exit("Timer expired")
}

func printProblem(p problem, num int) bool {
	fmt.Printf("Problem #%d: %s = ", num, p.q)
	var answer string
	fmt.Scanf("%s\n", &answer)
	return answer == p.a
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
