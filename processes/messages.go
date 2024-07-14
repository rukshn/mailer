package processes

import (
	"odk_mailer/models"
)

func CreateMessage(job models.Job, subject string, recipient string, content string) models.Message {
	db := models.ConnectDB()
	message := models.Message{JobID: job.ID, Job: job, Subject: subject, Recipient: recipient, Content: content, Status: false}
	db.Create(&message)
	return message
}

func BulkCreateMessage(messages []Message) bool {
	db := models.ConnectDB()
	db.Create(&messages)
	return true
}

func GetAllMessages() []models.Message {
	db := models.ConnectDB()
	var messages []models.Message
	db.Find(&messages)
	return messages
}

func GetMessagesByJobID(jobID int) []models.Message {
	db := models.ConnectDB()
	var messages []models.Message
	db.Where("job_id = ?", jobID).Find(&messages)
	return messages
}

func DeleteMessage(id int) bool {
	db := models.ConnectDB()
	var message models.Message
	db.Where("id = ?", id).First(&message)
	db.Delete(&message)
	return true
}

func UpdateMessage(id int, subject string, recipient string, content string, status bool) models.Message {
	db := models.ConnectDB()
	var message models.Message
	db.Where("id = ?", id).First(&message)
	message.Subject = subject
	message.Recipient = recipient
	message.Content = content
	message.Status = status
	db.Save(&message)
	return message
}

func ReadMessage(id int) models.Message {
	db := models.ConnectDB()
	var message models.Message
	db.Where("id = ?", id).First(&message)
	return message
}

func SendMessage(message models.Message) bool {
	return false
}
