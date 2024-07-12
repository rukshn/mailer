package parser

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadData() ([][]string, string) {
	if len(os.Args) != 3 {
		fmt.Println("Usage: parser <input data sheet> <message template>")
		os.Exit(1)
	}

	inputDataSheet := os.Args[1]
	messageTemplate := os.Args[2]

	if inputDataSheet == "" {
		fmt.Println("Input data sheet is required")
		os.Exit(1)
	}
	records := readCSV(inputDataSheet)

	if messageTemplate == "" {
		fmt.Println("Message template is required")
		os.Exit(1)
	}

	return records, messageTemplate
}

func readCSV(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

	readFile := csv.NewReader(file)
	records, err := readFile.ReadAll()
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	return records
}
