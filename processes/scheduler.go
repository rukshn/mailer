package processes

import (
	"odk_mailer/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RunPendingJobs() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var jobs []models.Job
	db.Where("status = ?", "pending").Find(&jobs)
	for _, job := range jobs {
		// Check if the job is due
		if job.Schedule.Before(time.Now()) {
			RunJob(job.Hash)
		}
	}
}
