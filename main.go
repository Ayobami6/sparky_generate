package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type loadingDoneMsg struct{}

// item type implementation for list item type interface
type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (itemDelegate) Height() int                               { return 1 }
func (itemDelegate) Spacing() int                              { return 0 }
func (itemDelegate) ShowSeparator() bool                       { return false }
func (itemDelegate) ShowSecondaryText() bool                   { return false }
func (itemDelegate) ShowPreview() bool                         { return false }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return SelectedItemStyle.Render("✔ " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

// tea model for state management
type model struct {
	projectList       list.Model
	spinner           spinner.Model
	projectChoice     string
	haveChosenProject bool
	quitting          bool
	done              bool
	showInput         bool
	textInput         textinput.Model
	inputString       string
	loading           bool
}

// model initialization constructor
func initialModel() model {
	// create spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = SpinnerStyle

	projects := []list.Item{
		item("Go"),
		item("Java SpringBoot"),
		item("NodeJS Nestjs"),
		item("Python Django"),
	}

	// create text input
	ti := textinput.New()
	ti.Placeholder = "Project Name"
	ti.PromptStyle = ProjectInputStyle
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	l := list.New(projects, itemDelegate{}, DefaultWidth, ListHeight)
	l.Title = "Choose a project type"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = TitleStyle
	l.Styles.PaginationStyle = PaginationStyle
	l.Styles.HelpStyle = HelpStyle

	return model{
		projectList:       l,
		spinner:           s,
		projectChoice:     "",
		haveChosenProject: false,
		quitting:          false,
		done:              false,
		showInput:         false,
		textInput:         ti,
		inputString:       "",
		loading:           false,
	}
}

// tea init command
func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

// UI Update to handle update changes in the terminal
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.projectList.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	case loadingDoneMsg:
		m.loading = false
		m.done = true
		return m, tea.Quit
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if !m.haveChosenProject {
		return updateProjectChoice(m, msg)
	}
	if m.showInput {
		return updateInput(m, msg)
	} else {
		var cmd tea.Cmd
		m.projectList, cmd = m.projectList.Update(msg)
		return m, cmd
	}

}

func (m model) View() string {
	var s string
	if m.quitting {
		return QuitTextStyle.Render("\nBye! See you Later!\n")
	}
	if !m.haveChosenProject {
		s = m.projectList.View()
	}
	if m.loading {
		return fmt.Sprintf("\n\n   %s Creating %s project...\n\n", m.spinner.View(), m.projectChoice)
	}
	if m.done {
		return doneView(m)
	}
	if m.showInput {
		return inputView(m)
	}

	return MainStyle.Render("\n" + s + "\n\n")
}

func updateProjectChoice(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() { //checking the keypress
		case "enter":
			m.haveChosenProject = true
			if idx, ok := m.projectList.SelectedItem().(item); ok {
				m.projectChoice = string(idx)
			}
			m.showInput = true
			return m, nil
		}
	}
	m.projectList, cmd = m.projectList.Update(msg)
	return m, cmd
}

func updateInput(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			m.inputString = m.textInput.Value()
			m.showInput = false
			m.loading = true
			return m, func() tea.Msg {
				// create a folder with the inputstring as name
				folderName := strings.ReplaceAll(m.inputString, " ", "_")
				generateProject(folderName, m.projectChoice)
				return loadingDoneMsg{}
			}
		}

	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
