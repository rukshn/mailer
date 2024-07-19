package output

import (
	"fmt"
	"odk_mailer/processes"

	"github.com/pterm/pterm"
)

func OutputJob(job []processes.Job, title string) {
	tableData := pterm.TableData{
		{"hash", "sender", "schedule", "status"},
	}

	for _, j := range job {
		tableData = append(tableData, []string{j.Hash, j.Sender, j.Schedule.String(), j.Status})
	}

	fmt.Println(title)
	pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(tableData).WithBoxed().Render()
}
