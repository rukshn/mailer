package input

import (
	"flag"
	"fmt"
	"odk_mailer/models"
	"odk_mailer/processes"
	"os"
	"time"

	"github.com/pterm/pterm"
)

type Input struct {
	Command string
	Job     processes.Job
	Value   []string
}

func ParseInput() Input {
	cmdNewJob := flag.NewFlagSet("new_job", flag.ExitOnError)
	newJobInputDataSheet := cmdNewJob.String("i", "", "Input data sheet")
	newJobMessageTemplate := cmdNewJob.String("t", "", "Message template")
	newJobSenderEmail := cmdNewJob.String("s", "", "Sender email")
	newJobSchedule := cmdNewJob.String("d", time.Now().String(), "Schedule")
	newJobHelp := cmdNewJob.Bool("h", false, "Print help message")

	cmdRunJob := flag.NewFlagSet("run", flag.ExitOnError)
	runJobHelp := cmdRunJob.Bool("h", false, "Print help message")
	runJobHash := cmdRunJob.String("i", "", "Job hash")

	cmdDeleteJob := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteJobHelp := cmdDeleteJob.Bool("h", false, "Print help message")
	deleteJobHash := cmdDeleteJob.String("i", "", "Job hash")

	cmdListJob := flag.NewFlagSet("list", flag.ExitOnError)
	listJobHelp := cmdListJob.Bool("h", false, "Print help message")

	switch os.Args[1] {

	// CREATE NEW JOB
	case "new":

		cmdNewJob.Parse(os.Args[2:])
		if len(os.Args) == 2 {
			pterm.DefaultBox.WithTitle(pterm.LightGreen("🆁 Mailer:: new")).Println("\n Usage: ./mailer" + pterm.FgYellow.Sprint("-new_job") + " -i <input data sheet> -t <message template> -s <sender email> -d <schedule> \n \n -i path to input data sheet (csv) file (required) \n -t path to message template file HTML or tmpl file (required) \n -s sender email (required) \n -d schedule time in RFC3339 format (optional) \n -h print help message")
			os.Exit(0)
		}
		if *newJobHelp {
			pterm.DefaultBox.WithTitle(pterm.LightGreen("🆁 Mailer:: new")).Println("\n Usage: ./mailer" + pterm.FgYellow.Sprint("-new_job") + " -i <input data sheet> -t <message template> -s <sender email> -d <schedule> \n \n -i path to input data sheet (csv) file (required) \n -t path to message template file HTML or tmpl file (required) \n -s sender email (required) \n -d schedule time in RFC3339 format (optional) \n -h print help message")
			os.Exit(0)
		}

		if *newJobMessageTemplate == "" {
			fmt.Println("Message template is required")
			os.Exit(1)
		}

		if *newJobInputDataSheet == "" {
			fmt.Println("Input data sheet is required")
			os.Exit(1)
		}

		if *newJobSenderEmail == "" {
			fmt.Println("Sender email is required")
			os.Exit(1)
		}

		var scheduleTime time.Time
		if *newJobSchedule == "" {
			scheduleTime = time.Now()
		} else {
			st, e := time.Parse(time.RFC3339, *newJobSchedule)
			scheduleTime = st
			if e != nil {
				panic(e)
			}
		}

		newJob := processes.Job{InputFile: *newJobInputDataSheet, TemplateFile: *newJobMessageTemplate, Sender: *newJobSenderEmail, Schedule: scheduleTime}
		return Input{Command: "new_job", Job: newJob}

	// RUN JOB
	case "run":
		cmdRunJob.Parse(os.Args[2:])
		if *runJobHelp {
			pterm.DefaultBox.WithTitle(pterm.LightGreen("🆁 Mailer:: run")).Println("\n Usage: ./mailer" + pterm.FgYellow.Sprint("-run") + " -i <job hash> \n \n -i job hash (required) \n -h print help message")
		}

		return Input{Command: "run_job", Job: processes.Job{Hash: *runJobHash}}

	// DELETE JOB
	case "delete":
		cmdDeleteJob.Parse(os.Args[2:])
		if *deleteJobHelp {
			pterm.DefaultBox.WithTitle(pterm.LightGreen("🆁 Mailer:: delete")).Println("\n Usage: ./mailer" + pterm.FgYellow.Sprint("-delete") + " -i <job hash> \n \n -i job hash (required) \n -h print help message")
		}
		return Input{Command: "delete_job", Job: processes.Job{Hash: *deleteJobHash}}

	case "list":
		cmdListJob.Parse(os.Args[2:])
		if *listJobHelp {
			pterm.DefaultBox.WithTitle(pterm.LightGreen("🆁 Mailer:: list")).Println("\n Usage: ./mailer" + pterm.FgYellow.Sprint("-list") + " \n \n -h print help message")
		}
		return Input{Command: "list_jobs"}
	default:
		fmt.Println(os.Args[1])
	}

	return Input{Command: "help"}
}

func initDb() {
	models.Migrate()
}
