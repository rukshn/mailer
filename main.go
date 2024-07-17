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
		newJob := processes.Job(parseCommand.Job)
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
		fmt.Println("Running job", parseCommand.Job.Hash)
		processes.RunJob(parseCommand.Job.Hash)
	}

	if parseCommand.Command == "delete_job" {
		fmt.Println("Deleting job", parseCommand.Job.Hash)
		processes.DeleteJob(parseCommand.Job.Hash)
	}

	if parseCommand.Command == "run_schedules" {
		fmt.Println("Running schedules")
		processes.RunPendingJobs()
	}
}
