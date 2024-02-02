package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
	os.Setenv("LC_CTYPE", "en_US.UTF-8")
	p := tea.NewProgram(
		NewSimplePage(),
		tea.WithAltScreen(),
	)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
