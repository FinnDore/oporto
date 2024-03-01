package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type UnselectedServicesModel struct {
	allServices []string
}

func NewUnselectedServices() UnselectedServicesModel {
	return UnselectedServicesModel{}
}

func (s UnselectedServicesModel) Init() tea.Cmd { return nil }

func (s UnselectedServicesModel) View() string {
	if s.allServices == nil {
		return ""
	}
	return strings.Join(s.allServices, "\n")
}
func (s *UnselectedServicesModel) SetServices(services []string) {
	s.allServices = services
}

func (f UnselectedServicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	}
	return f, nil
}
