package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

type QuestionReader interface {
	ParseQuestions(r io.Reader) ([]Question, error)
}

type CsvReader struct {
}

func (c CsvReader) ParseQuestions(r io.Reader) ([]Question, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var out []Question
	for _, record := range records {
		out = append(out, Question{record[0], record[1]})
	}
	return out, err
}

type TxtReader struct {
}

func (t TxtReader) ParseQuestions(r io.Reader) ([]Question, error) {
	scanner := bufio.NewScanner(r)
	var i int
	var question string
	var answer string
	var out []Question
	for scanner.Scan() {
		if i%2 == 0 {
			question = scanner.Text()
		} else {
			answer = scanner.Text()
			out = append(out, Question{question, answer})
		}
		i++
	}
	return out, nil
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
	return out
}

var (
	f = flag.String("f", "", "-f path to file")
)

func main() {
	flag.Parse()
	var r QuestionReader
	if (*f)[len(*f)-3:] == "csv" {
		r = CsvReader{}
	} else if (*f)[len(*f)-3:] == "txt" {
		r = TxtReader{}
	}

	file, err := os.Open(*f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0

	questions, _ := r.ParseQuestions(file)

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
