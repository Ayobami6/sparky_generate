package main

import "fmt"

func doneView(m model) string {
	return QuitTextStyle.Render("Your Project " + m.inputString + " is ready!")
}

func projectListView(m model) string {
	return m.projectList.View()
}

func inputView(m model) string {
	return fmt.Sprintf("Enter your project name: %s", m.textInput.View()+"\n")
}
