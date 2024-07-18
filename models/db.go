package models

import (
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Schedule time.Time
	Hash     string
	Status   string
	Sender   string
	ID       int
}

type Message struct {
	Content   string `gorm:"type:text"`
	Subject   string
	Recipient string
	Job       Job
	Status    bool
	ID        int
	JobID     int
}

type Settings struct {
	gorm.Model
	Key   string
	Value string
}

func Migrate() {
	dbPath, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/test.db")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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
