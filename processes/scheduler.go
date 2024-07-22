package processes

import (
	"odk_mailer/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RunPendingJobs() []Job {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var jobs []models.Job
	var maskJobs []Job
	db.Where("status = ?", "pending").Find(&jobs)
	for _, job := range jobs {
		// Check if the job is due
		if job.Schedule.Before(time.Now()) {
			completed_job := RunJob(job.Hash)
			maskJobs = append(maskJobs, Job{Schedule: completed_job.Schedule, Status: completed_job.Status, Hash: completed_job.Hash, Sender: completed_job.Sender})
		}
	}
	return maskJobs
}
