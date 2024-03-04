package main

import (
	table "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

var unfocusedTableStyle1 = table.Styles{
	Selected: lipgloss.NewStyle(),
	Header:   table.DefaultStyles().Header,
	Cell:     table.DefaultStyles().Cell,
}

type SelectedServicesModel struct {
	unselectedServices []string
	table              table.Model
}

func NewSelectedServices() SelectedServicesModel {
	unselectedServices := SelectedServicesModel{
		table: table.New(
			table.WithColumns([]table.Column{{Title: "Name", Width: 20}}),
			table.WithRows([]table.Row{}),
			table.WithFocused(true),
			table.WithHeight(7),
		),
	}

	unselectedServices.table.SetStyles(unfocusedTableStyle)
	return unselectedServices

}

func (s SelectedServicesModel) Init() tea.Cmd { return nil }

func (s SelectedServicesModel) View() string {
	if s.unselectedServices == nil {
		return ""
	}
	return s.table.View()
}
func (s *SelectedServicesModel) SetServices(services []string) {
	rows := Map(services, func(service string) table.Row {
		return table.Row{service}
	})
	s.table.SetRows(rows)
	s.unselectedServices = services
}

func (f *SelectedServicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "h", "left":
			f.table.Blur()
			f.table.SetStyles(unfocusedTableStyle1)
		case "l", "right":
			f.table.SetStyles(table.DefaultStyles())
			f.table.Focus()
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
