package main

import (
	"fmt"
	"odk_mailer/output"
	"odk_mailer/parser"
	"odk_mailer/processes"
)

func main() {
	parseCommand := parser.ParseInput()
	fmt.Println(parseCommand)
	if parseCommand.Command == "new_job" {
		fmt.Println("Creating new job")
		newJob := processes.NewJob(parseCommand.NewJob)
		generatedNewJob := processes.GenerateNewJob(newJob)
		fmt.Println(generatedNewJob)
	}

	if parseCommand.Command == "list_jobs" {
		listJobs := processes.GetAllJobs()
		output.GenerateAllJobTable(listJobs)
	}

	if parseCommand.Command == "list_messages" {
		fmt.Println(processes.GetAllMessages())
	}

	if parseCommand.Command == "run_job" {
		fmt.Println("Running job", parseCommand.NewJob.Hash)
		processes.RunJob(parseCommand.NewJob.Hash)
	}
}
