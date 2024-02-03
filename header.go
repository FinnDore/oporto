package main

import (
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type headerModel struct {
}

var mainHeaerStyle = lipgloss.NewStyle().Inline(true).Bold(true).Foreground(lipgloss.Color("#5900ff")).MaxHeight(1)

var headerAltStyle = lipgloss.NewStyle().Inline(true).
	Foreground(lipgloss.Color("#828181")).MaxHeight(1)

func NewheaderModel() headerModel {
	return headerModel{}
}

func (s headerModel) Init() tea.Cmd { return nil }

func (s headerModel) View() string {
	return headerStyle.Render("OPORTO") + headerAltStyle.Render(" - v0.0.1-"+Commit)
}

func (s headerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}
