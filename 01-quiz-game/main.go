package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Quiz Game")

	filename := flag.String("f", "questions.csv", "Specify questions filename")
	timeForQuestion := flag.Int("t", 30, "Specify time for question")
	flag.Parse()

	questions, total := loadFile(*filename)
	correctAnswers := playGame(questions, *timeForQuestion)

	fmt.Printf("You've answered correct to %d of %d questions\n", correctAnswers, total)
}

func loadFile(filename string) ([][]string, int){
	fmt.Printf("Loading file %s\n", filename)
	fileContent, loadFileError := os.ReadFile(filename)
	if loadFileError != nil {
		fmt.Println("Error occured during loading file:", loadFileError)
		panic(loadFileError)
	}

	reader := csv.NewReader(strings.NewReader(string(fileContent)))

	questions, readCsvError := reader.ReadAll()

	if readCsvError!= nil {
		fmt.Println("Error occured during reading csv values:", readCsvError)
		panic(readCsvError)
	}

	fmt.Printf("Loaded %d questions\n", len(questions))
	return questions, len(questions)
}

func playGame(questions [][]string, timeForQuestion int) (int) {
	correct := 0
	currentQuestion := 1
	for _, question := range questions {
		fmt.Printf("Current question: %d\n", currentQuestion)
		fmt.Printf("%s [%s]\n> ", question[0], question[1])
		answer := ""
		fmt.Scanln(&answer)
		if strings.EqualFold(answer, question[1]) {
			fmt.Println("Correct")
			correct++
		} else {
			fmt.Println("Wrong")
		}
		fmt.Println()
		currentQuestion++
	}
	return correct
}