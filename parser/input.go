package parser

import (
	"flag"
	"fmt"
	"odk_mailer/models"
	"odk_mailer/processes"
	"os"
	"time"
)

type Input struct {
	Command string
	Job     processes.Job
	Value   []string
}

func ParseInput() Input {
	cmd := flag.String("c", "", "Command")
	help := flag.String("h", "", "Help")
	inputDataSheet := flag.String("i", "", "Input data sheet")
	messageTemplate := flag.String("t", "", "Message template")
	senderEmail := flag.String("s", "", "Sender email")
	schedule := flag.String("d", time.Now().String(), "Schedule")
	init := flag.Bool("init", false, "Initialize database")
	jobHash := flag.String("j", "", "Job hash")
	runSchedules := flag.Bool("rs", false, "Run schedules")

	flag.Parse()

	if *cmd == "h" || *cmd == "help" || *help != "" {
		fmt.Println("Usage: ./mailer -c new_job -i <input data sheet> -t <message template> -s <sender email> -d <schedule>")
		os.Exit(0)
	}

	if *runSchedules {
		return Input{Command: "run_schedules"}
	}

	if *init {
		initDb()
		fmt.Println("Database initialized")
		os.Exit(0)
	}

	if *jobHash != "" {
		if *cmd == "run_jon" {
			return Input{Command: "read_job", Job: processes.Job{Hash: *jobHash}}
		}

		if *cmd == "delete_job" {
			return Input{Command: "delete_job", Job: processes.Job{Hash: *jobHash}}
		}
	}

	if *cmd == "new_job" {
		if *messageTemplate == "" {
			fmt.Println("Message template is required")
			os.Exit(1)
		}

		if *inputDataSheet == "" {
			fmt.Println("Input data sheet is required")
			os.Exit(1)
		}

		if *senderEmail == "" {
			fmt.Println("Sender email is required")
			os.Exit(1)
		}

		var scheduleTime time.Time
		if *schedule == "" {
			scheduleTime = time.Now()
		} else {
			st, e := time.Parse(time.RFC3339, *schedule)
			scheduleTime = st
			if e != nil {
				panic(e)
			}
		}
		newJob := processes.Job{InputFile: *inputDataSheet, TemplateFile: *messageTemplate, Sender: *senderEmail, Schedule: scheduleTime}
		return Input{Command: *cmd, Job: newJob}
	}

	if *cmd == "list_jobs" {
		return Input{Command: *cmd}
	}

	if *cmd == "list_messages" {
		return Input{Command: *cmd}
	}

	return Input{Command: *cmd}
}

func initDb() {
	models.Migrate()
}
