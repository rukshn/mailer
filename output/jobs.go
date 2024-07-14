package output

import (
	"odk_mailer/models"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GenerateAllJobTable(jobs []models.Job) {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true).
		SetBordersColor(tcell.ColorBlue)
	headers := []string{"ID", "Sender", "Hash", "Status"}
	for i, header := range headers {
		table.SetCell(0, i,
			tview.NewTableCell(header).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
	}
	for i, job := range jobs {
		table.SetCell(i+1, 0,
			tview.NewTableCell(strconv.Itoa(job.ID)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
		table.SetCell(i+1, 1,
			tview.NewTableCell(job.Sender).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
		table.SetCell(i+1, 2,
			tview.NewTableCell(job.Hash).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
		table.SetCell(i+1, 3,
			tview.NewTableCell(job.Status).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter).
				SetSelectable(false))
	}
	flex := tview.NewFlex().
		AddItem(table, 0, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
