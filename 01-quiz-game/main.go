package main

//todo: add parameter filename
//todo: calculate valid answers and total questions
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Quiz Game")

	var filename string
	flag.StringVar(&filename, "f", "questions.csv", "Specify questions filename")
	flag.Parse()

	questions, total := loadFile(filename)
	correctAnswers := playGame(questions)

	fmt.Printf("You've answered correct to %d of %d questions\n", correctAnswers, total)
}

func loadFile(filename string) ([][]string, int){
	fmt.Printf("Loading file %s\n", filename)
	fileContent, _ := os.ReadFile(filename)
	reader := csv.NewReader(strings.NewReader(string(fileContent)))

	questions, _ := reader.ReadAll()
	fmt.Printf("Loaded %d questions\n", len(questions))
	return questions, len(questions)
}

func playGame(questions [][]string) (int) {
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