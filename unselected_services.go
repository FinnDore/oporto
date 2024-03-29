package main

import (
	table "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

var unfocusedTableStyle = table.Styles{
	Selected: lipgloss.NewStyle(),
	Header:   table.DefaultStyles().Header,
	Cell:     table.DefaultStyles().Cell,
}

type UnselectedServicesModel struct {
	unselectedServices []string
	table              table.Model
}

func NewUnselectedServices() UnselectedServicesModel {
	unselectedServices := UnselectedServicesModel{
		table: table.New(
			table.WithColumns([]table.Column{{Title: "Name", Width: 20}}),
			table.WithRows([]table.Row{}),
			table.WithFocused(true),
			table.WithHeight(7),
		),
	}

	return unselectedServices

}

func (s UnselectedServicesModel) Init() tea.Cmd { return nil }

func (s UnselectedServicesModel) View() string {
	if s.unselectedServices == nil {
		return ""
	}
	return s.table.View()
}
func (s *UnselectedServicesModel) SetServices(services []string) {
	rows := Map(services, func(service string) table.Row {
		return table.Row{service}
	})
	s.table.Focus()
	s.table.SetRows(rows)
	s.unselectedServices = services
}

func (f *UnselectedServicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "h", "left":
			f.table.Focus()
			f.table.SetStyles(table.DefaultStyles())
		case "l", "right":
			f.table.SetStyles(unfocusedTableStyle)
			f.table.Blur()
		case "up", "k":
			if f.table.Focused() {
				f.table.MoveUp(0)
			}
		case "down", "j":
			if f.table.Focused() {
				f.table.MoveDown(0)
			}
		}
	}

	var cmd tea.Cmd
	f.table, cmd = f.table.Update(msg)
	return f, cmd
}
