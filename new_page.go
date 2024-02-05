package main

import (
	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

// MODEL DATA
type simplePage struct {
	flexBox     *stickers.FlexBox
	header      headerModel
	footer      footerModel
	config      Config
	renderCount int
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

	s := simplePage{
		flexBox:     stickers.NewFlexBox(0, 0).SetStyle(pageStyle),
		header:      NewheaderModel(),
		footer:      NewFooter(),
		config:      config,
		renderCount: 0,
	}

	header := s.flexBox.NewRow().AddCells(
		[]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(0, 0).SetContent(s.header.View()).SetStyle(headerStyle),
		},
	)

	services, err := GetServices()

	if err != nil {
		panic(nil)
	}
	var thing = ""
	for _, service := range services.Application {
		thing += blockText.Render(service.Name) + "\n"
	}

	mainContent := s.flexBox.NewRow().AddCells(
		[]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(1, 100).
				SetContent(thing).SetStyle(mainStyle),
			stickers.NewFlexBoxCell(1, 100).SetStyle(mainStyle),
		},
	)

	footer := s.flexBox.NewRow().AddCells(
		[]*stickers.FlexBoxCell{
			stickers.NewFlexBoxCell(0, 5).SetContent(s.footer.View()).SetStyle(footerStyle),
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
	s.renderCount++

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
