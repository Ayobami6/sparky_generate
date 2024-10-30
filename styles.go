package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const ListHeight = 14

const DefaultWidth = 20

// lipglos styles for terminal UI
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFDE59")).MarginLeft(2)
	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("150")).PaddingLeft(2)
	PaginationStyle = list.DefaultStyles().PaginationStyle.
			PaddingLeft(4)
	HelpStyle = list.DefaultStyles().HelpStyle.
			PaddingLeft(4)
	QuitTextStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFDE59")).Margin(2, 0, 2, 4)
	MainStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFDE59")).MarginLeft(2)
	SpinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("120")).MarginLeft(2)
	ProjectInputStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFDE59")).Bold(true)
	InputTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFDE59")).MarginLeft(2)
)
