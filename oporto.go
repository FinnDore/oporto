package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		NewSimplePage("This app is under construction"),
		tea.WithAltScreen(),
	)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
