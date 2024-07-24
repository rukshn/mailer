package output

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func LogToFile(logId string, logMessage string) {
	logPath, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/logs/" + logId + ".log")
	fmt.Println(logPath)
	if err != nil {
		fmt.Println("Error getting log path: ", err)
	}

	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(logMessage)
}

func ReadLog(logId string) {
	logPath, err := filepath.Abs(filepath.Dir(os.Args[0]) + "logs/" + logId + ".log")
	if err != nil {
		fmt.Println("Error getting log path: ", err)
	}

	f, err := os.OpenFile(logPath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	fmt.Println("📜 Reading log file: ", logPath)
	fmt.Println(f)
}
