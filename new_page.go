package main

import (
	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

// MODEL DATA
type simplePage struct {
	flexBox *stickers.FlexBox
	header  headerModel
	footer  footerModel
}

var pageStyle = lipgloss.NewStyle().Margin(1)

var mainStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFFFFF0F"))

var headerStyle = lipgloss.NewStyle().
	Height(1).Bold(true).Foreground(lipgloss.Color("#5900ff"))

var footerStyle = lipgloss.NewStyle().Height(10)

func NewSimplePage() simplePage {
	s := simplePage{
		flexBox: stickers.NewFlexBox(0, 0).SetStyle(pageStyle),
		header:  NewheaderModel(),
		footer:  NewFooter(),
	}

	header := s.flexBox.NewRow().AddCells(
		[]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(0, 0).SetContent(s.header.View()).SetStyle(headerStyle),
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
			stickers.NewFlexBoxCell(0, 1).SetContent(s.footer.View()).SetStyle(footerStyle),
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
