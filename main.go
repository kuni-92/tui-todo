package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	listView := tview.NewTextView()
	listView.SetTitle("ToDo").SetBorder(true)

	textBox := tview.NewInputField()
	textBox.SetLabel("ToDo: ").SetTitle("New ToDo").SetBorder(true)

	textBox.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			listView.SetText(listView.GetText(true) + textBox.GetText() + "\n")
			textBox.SetText("")
			return nil
		}
		return event
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow).
		AddItem(textBox, 4, 0, true).
		AddItem(listView, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
