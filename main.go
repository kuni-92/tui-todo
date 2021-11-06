package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const FILEPATH = "ToDoFile"

func main() {
	todos, err := readFile()
	if err != nil {
		fmt.Println(err)
	}

	app := tview.NewApplication()

	listView := tview.NewTextView()
	listView.SetTitle("ToDo").SetBorder(true)
	listView.SetText(todos)

	textBox := tview.NewInputField()
	textBox.SetLabel("ToDo: ").SetTitle("New ToDo").SetBorder(true)

	textBox.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			todoText := textBox.GetText()
			if todoText == "" {
				return nil
			}

			if err := writeFile(todoText); err != nil {
				fmt.Println(err)
			}

			listView.SetText(listView.GetText(true) + todoText + "\n")
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

// readFile reads ToDo list from the file.
func readFile() (string, error) {
	if _, err := os.Stat(FILEPATH); os.IsNotExist(err) {
		return "", nil
	}

	f, err := os.OpenFile(FILEPATH, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	todos := string(buf)
	return todos, nil
}

// writeFile writes ToDo to the file.
func writeFile(todo string) error {
	f, err := os.OpenFile(FILEPATH, (os.O_RDWR | os.O_CREATE | os.O_APPEND), 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, todo+"\n"); err != nil {
		return err
	}
	return nil
}
