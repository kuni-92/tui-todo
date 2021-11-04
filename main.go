package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("New Task"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("ToDo List"), 0, 3, false), 0, 2, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
