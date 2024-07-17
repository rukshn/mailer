package processes

import (
	"bytes"
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
	headers := records[0]
	temp, err := template.New(templateStr).ParseFiles(templateStr)
	if err != nil {
		DeleteJob("", JobID)
		panic(err)
	}

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
