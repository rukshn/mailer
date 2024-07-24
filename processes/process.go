package processes

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
)

type Message struct {
	Recipient string
	Subject   string
	Content   string
	JobID     int
	Status    bool
}

func processTemplate(temp *template.Template, data map[string]string) string {
	var tpl bytes.Buffer

	temp.Execute(&tpl, data)
	return tpl.String()
}

func ProcessRecords(records [][]string, templateStr string, JobID int, Fields []string) []Message {
	templatePath, err := filepath.Abs(templateStr)
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles(templatePath)
	temp.Option("missingkey=zero")
	if err != nil {
		DeleteJob("", JobID)
		panic(err)
	}

	headerToValueMap := make(map[string]string)

	var messages []Message
	fullHeaders := records[0]
	fullHeaderList := map[string]int{}

	for i, fhi := range fullHeaders {
		fullHeaderList[fhi] = i
	}

	var headers []string
	if len(Fields) > 1 {
		headers = Fields
		headers = append(headers, "email")
	} else if len(Fields) == 1 && Fields[0] == "" {
		headers = fullHeaders
	}

	fmt.Println("Headers", headers)
	for _, record := range records[1:] {
		for _, header := range headers {
			headerToValueMap[header] = record[fullHeaderList[header]]
		}

		processedMessage := processTemplate(temp, headerToValueMap)
		splitMessage := strings.Split(processedMessage, "%%--%%--%%")
		msg := Message{Recipient: headerToValueMap["email"], Content: splitMessage[1], JobID: JobID, Status: false, Subject: splitMessage[0]}

		messages = append(messages, msg)

	}
	return messages
}
