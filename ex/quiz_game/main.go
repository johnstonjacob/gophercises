package main

import (
	"encoding/csv"
	//	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting quiz...")
	qa := getQuestions()

	gameLoop(qa)
}

func getQuestions() [][]string {
	path := "/Users/jacobjohnston/Go/src/github.com/johnstonjacob/gophercises/ex/quiz_game/problems.csv"
	osReader, _ := os.Open(path)
	csvReader := csv.NewReader(osReader)

	qa, _ := csvReader.ReadAll()

	return qa
}

func getInput() string {
	fmt.Print("(Answer)  ")
	var input string
	fmt.Scanln(&input)

	return input
}

func gameLoop(qa [][]string) {
	score := 0

	for _, tuple := range qa {
		q := tuple[0]
		a := tuple[1]

		fmt.Println(q)

		input := getInput()

		if input == a {
			score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect..")
		}
	}

	fmt.Printf("Final score: %d / %d\n", score, len(qa))
}
