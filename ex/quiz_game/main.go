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

type game struct {
	score     int
	questions int
}

func main() {
	var shuffle bool
	var timer int

	game := game{0, 0}

	flag.BoolVar(&shuffle, "s", false, "Shuffle the questions. Default: false")
	flag.IntVar(&timer, "t", 90, "Sets time for timer in seconds. Default: 90")

	flag.Parse()

	fmt.Println("Starting quiz...")
	qa := getQuestions(shuffle)

	go quizTimer(timer, &game)
	gameLoop(qa, &game)
}

func quizTimer(t int, g *game) {
	d := time.Duration(t) * time.Second
	time.AfterFunc(d, func() {
		fmt.Println("\nTimes up!")
		fmt.Printf("Final score: %d / %d\n", g.score, g.questions)
		os.Exit(0)
	})

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

func gameLoop(qa [][]string, g *game) {
	g.questions = len(qa)

	for _, tuple := range qa {
		q, a := tuple[0], tuple[1]

		fmt.Println(q)

		input := getInput()

		if compareString(a, input) {
			g.score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect..")
		}
	}

	fmt.Printf("Final score: %d / %d\n", g.score, len(qa))
}
