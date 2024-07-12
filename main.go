package main

import (
	"fmt"
	"odk_mailer/parser"
)

func main() {
	records, message := parser.ReadData()
	fmt.Println(records)
	fmt.Println(message)
}
