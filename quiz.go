package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type QuestionReader interface {
	ParseQuestions(r io.Reader) ([]Question, error)
}

type Question struct {
	question string
	answer   string
}

func readCsv(filename string) []Question {
	csvFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var out []Question
	for _, record := range records {
		out = append(out, Question{record[0], record[1]})
	}
	//fmt.Print(out)
	// Допишите код здесь
	return out
}

func main() {
	total := 0

	questions := readCsv("problems.csv")
	// Пройтись циклом. Вывести вопрос, предложить пользователю ввести ответ.
	// Если ответ правильный, увеличить total.
	// for
	for _, question := range questions {
		var result string
		fmt.Printf("%s: ", question.question)
		_, _ = fmt.Scanf("%s", &result)
		if question.answer == result {
			total++
		}
	}
	fmt.Printf("You got %d / %d", total, len(questions))
}
