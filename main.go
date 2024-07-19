package main

import (
	"fmt"
	"odk_mailer/input"
	"odk_mailer/output"
	"odk_mailer/processes"
)

func main() {
	parseCommand := input.ParseInput()

	if parseCommand.Command == "new_job" {
		newJob := processes.Job(parseCommand.Job)
		generatedNewJob := processes.GenerateNewJob(newJob)
		output.OutputJob(generatedNewJob, "✅ New job created")
	}

	if parseCommand.Command == "list_messages" {
		listMessages := processes.GetAllMessages()
		output.GenerateAllMessagesTable(listMessages)
	}

	if parseCommand.Command == "list_jobs" {
		listJobs := processes.GetAllJobs()
		output.OutputJob(listJobs, "📋 All jobs")
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
