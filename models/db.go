package models

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID       int
	Hash     string
	Status   string
	Schedule time.Time
	Sender   string
}

type Message struct {
	ID        int
	JobID     int
	Job       Job
	Subject   string
	Recipient string
	Content   string `gorm:"type:text"`
	Status    bool
}

type Settings struct {
	gorm.Model
	Key   string
	Value string
}

func Migrate() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Job{})
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&Settings{})
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
