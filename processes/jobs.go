package processes

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"odk_mailer/models"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Job struct {
	InputFile    string
	TemplateFile string
	Sender       string
	Schedule     time.Time
	JobID        int
	Hash         string
}

func ReadJob(hash string) models.Job {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var job models.Job
	db.Where("hash = ?", hash).First(&job)
	return job
}

func createJob(schedule time.Time, sender string) models.Job {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		panic("failed to generate random bytes")
	}

	hashBytes := sha256.Sum256(randomBytes)
	hashHex, err := hex.EncodeToString(hashBytes[:]), nil

	job := models.Job{Hash: hashHex, Schedule: schedule, Status: "pending", Sender: sender}
	db.Create(&job)
	return job
}

func UpdateJob(job models.Job) models.Job {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Save(&job)
	return job
}

func DeleteJob(hash string) bool {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var job models.Job
	var messages []models.Message
	db.Where("hash = ?", hash).First(&job)
	db.Where("job_id = ?", job.ID).Find(&messages)
	db.Delete(&messages)
	db.Delete(&job)
	return true
}

func RunJob(hash string) {
	job := ReadJob(hash)
	job.Status = "processing"
	UpdateJob(job)
	messages := GetMessagesByJobID(job.ID)

	for _, message := range messages {
		mail := SendMail(message, job)
		fmt.Println("Mail sent: ", mail)
		message.Status = true
		UpdateMessage(message.ID, message.Subject, message.Recipient, message.Content, message.Status)
	}
	job.Status = "completed"
	UpdateJob(job)
}

func GenerateNewJob(job Job) Job {
	newJob := createJob(job.Schedule, job.Sender)
	job.JobID = newJob.ID
	inputData := readCSV(job.InputFile)
	fmt.Println(inputData)
	messages := ProcessRecords(inputData, job.TemplateFile, job.JobID)
	BulkCreateMessage(messages)
	return job
}

func GetAllJobs() []models.Job {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var jobs []models.Job
	db.Find(&jobs)
	return jobs
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
