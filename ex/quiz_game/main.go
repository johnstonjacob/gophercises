package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var shuffle bool
	flag.BoolVar(&shuffle, "s", false, "Shuffle the questions. Default: false")

	flag.Parse()

	fmt.Println("Starting quiz...")
	qa := getQuestions(shuffle)

	gameLoop(qa)
}

func getQuestions(shuffle bool) [][]string {
	path := "/Users/jacobjohnston/Go/src/github.com/johnstonjacob/gophercises/ex/quiz_game/problems.csv"
	osReader, _ := os.Open(path)
	csvReader := csv.NewReader(osReader)

	qa, _ := csvReader.ReadAll()

	if shuffle {
		return shuffleQuestions(qa)
	}

	return qa
}

func shuffleQuestions(qa [][]string) [][]string {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Shuffling questions..")
	rand.Shuffle(len(qa), func(i, j int) {
		qa[i], qa[j] = qa[j], qa[i]
	})

	return qa
}

func getInput() string {
	fmt.Print("(Answer)  ")
	var input string
	fmt.Scanln(&input)

	return input
}

func compareString(a string, i string) bool {
	a = strings.Trim(a, " ")
	a = strings.ToLower(a)

	i = strings.Trim(i, " ")
	i = strings.ToLower(i)

	return a == i
}

func gameLoop(qa [][]string) {
	score := 0

	for _, tuple := range qa {
		q, a := tuple[0], tuple[1]

		fmt.Println(q)

		input := getInput()

		if compareString(a, input) {
			score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect..")
		}
	}

	fmt.Printf("Final score: %d / %d\n", score, len(qa))
}
