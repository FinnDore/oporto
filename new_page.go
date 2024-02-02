package main

import (
	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

// MODEL DATA
type simplePage struct {
	flexBox *stickers.FlexBox
}

var mainStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFFFFF0F")).Margin(1)

var headerStyle = lipgloss.NewStyle().
	Height(10)

var footerStyle = lipgloss.NewStyle()

func NewSimplePage(text string) simplePage {
	s := simplePage{
		flexBox: stickers.NewFlexBox(0, 0),
	}

	header := s.flexBox.NewRow().AddCells(
		[]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(1, 1).SetStyle(headerStyle),
		},
	)

	mainContent := s.flexBox.NewRow().AddCells(
		[]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(16, 9).SetStyle(mainStyle),
			stickers.NewFlexBoxCell(16, 9).SetStyle(mainStyle),
		},
	)

	footer := s.flexBox.NewRow().AddCells(
		[]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(1, 1).SetStyle(footerStyle),
		},
	)

	s.flexBox.AddRows([]*stickers.FlexBoxRow{header})
	s.flexBox.AddRows([]*stickers.FlexBoxRow{mainContent})
	s.flexBox.AddRows([]*stickers.FlexBoxRow{footer})
	return s
}

func (s simplePage) Init() tea.Cmd { return nil }

// VIEW
func (s simplePage) View() string {
	return s.flexBox.Render()
}

// UPDATE
func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.WindowSizeMsg:
		s.flexBox.SetWidth(msg.(tea.WindowSizeMsg).Width)
		s.flexBox.SetHeight(msg.(tea.WindowSizeMsg).Height)
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
		case "q":
			return s, tea.Quit
		}
	}
	return s, nil
}
