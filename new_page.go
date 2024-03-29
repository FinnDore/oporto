package main

import (
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

// MODEL DATA
type simplePage struct {
	header             headerModel
	footer             footerModel
	unselectedServices UnselectedServicesModel
	selectedServices   SelectedServicesModel
	config             Config
	applications       Applications
}

var pageStyle = lipgloss.NewStyle().Margin(1)
var mainStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFFFFF0F"))
var headerStyle = lipgloss.NewStyle().
	Bold(true).Foreground(lipgloss.Color("#5900ff"))
var footerStyle = lipgloss.NewStyle()
var blockText = lipgloss.NewStyle()

func NewSimplePage() simplePage {
	config, _ := LoadConfig()

	aplications, err := GetServices()

	if err != nil {
		panic(err)
	}

	s := simplePage{
		header:             NewheaderModel(),
		unselectedServices: NewUnselectedServices(),
		selectedServices:   NewSelectedServices(),
		footer:             NewFooter(),
		config:             config,
	}

	services := Map(aplications.Application, func(app Aplication) string {
		return app.Name
	})

	s.unselectedServices.SetServices(services)
	s.selectedServices.SetServices(services)
	return s
}

func (s simplePage) Init() tea.Cmd { return nil }

// VIEW
func (s simplePage) View() string {
	return s.header.View() + "\n" + s.unselectedServices.View() + s.selectedServices.View() + "\n" + s.footer.View()
}

// UPDATE
func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.WindowSizeMsg:
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
		case "q":
			return s, tea.Quit
		}
	}

	var cmds []tea.Cmd
	_, cmd1 := s.unselectedServices.Update(msg)

	_, cmd2 := s.selectedServices.Update(msg)
	cmds = append(cmds, cmd1, cmd2)
	return s, tea.Batch(cmds...)
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
