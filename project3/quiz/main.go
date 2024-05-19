package main

import (
	"fmt"
	"os"
	"quiz/game"
	"quiz/questions"
	"quiz/shuffler"
)

func main() {

	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Printf("error reading question file: %w", err)
		os.Exit(1)
	}

	shuffler.Shuffle(questions)

	correctAnswer := game.Run(questions)

	fmt.Printf("Correct answers: %d", correctAnswer)

}
