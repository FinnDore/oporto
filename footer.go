package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type footerModel struct {
	help helpModel
}

func NewFooter() footerModel {
	return footerModel{
		help: NewHelpSection(),
	}
}

func (s footerModel) Init() tea.Cmd { return nil }

func (f footerModel) View() string {
	return f.help.View()
}

func (f footerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	}
	return f, nil
}
