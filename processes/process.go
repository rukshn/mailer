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

func ProcessRecords(records [][]string, templateStr string, JobID int) []Message {
	templatePath, err := filepath.Abs(templateStr)
	if err != nil {
		panic(err)
	}

	headers := records[0]
	temp, err := template.ParseFiles(templatePath)
	if err != nil {
		DeleteJob("", JobID)
		panic(err)
	}

	fmt.Println(temp.Root.Nodes[1].String())

	headerToValueMap := make(map[string]string)

	var messages []Message
	for _, record := range records[1:] {

		for i, header := range headers {
			headerToValueMap[header] = record[i]
		}

		processedMessage := processTemplate(temp, headerToValueMap)
		splitMessage := strings.Split(processedMessage, "%%--%%--%%")
		msg := Message{Recipient: headerToValueMap["email"], Content: splitMessage[1], JobID: JobID, Status: false, Subject: splitMessage[0]}

		messages = append(messages, msg)

	}
	return messages
}
