package main

//todo: add parameter filename
//todo: calculate valid answers and total questions
import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

var filename string = "questions.csv"

func main() {
	fmt.Println("Quiz Game")
	questions, total := loadFile(filename)

}

func loadFile(filename string) ([][]string, int){
	fmt.Printf("Loading file %s\n", filename)
	fileContent, _ := os.ReadFile(filename)
	reader := csv.NewReader(strings.NewReader(string(fileContent)))

	questions, _ := reader.ReadAll()
	fmt.Printf("Loaded %d questions\n", len(questions))
	return questions, len(questions)
}