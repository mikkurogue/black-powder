package components

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Model struct {
	table *table.Table
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.table = m.table.Width(msg.Width)
		m.table = m.table.Height(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
		}
	}
	return m, cmd
}

func (m Model) View() string {
	return "\n" + m.table.String() + "\n"
}

func RenderTable() {
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
	selectedStyle := baseStyle.Foreground(lipgloss.Color("#01BE85")).Background(lipgloss.Color("#00432F"))
	typeColors := map[string]lipgloss.Color{
		"Go":         lipgloss.Color("#90d1ff"),
		"TypeScript": lipgloss.Color("#7D5AFC"),
	}
	headers := []string{"#", "PROJECT", "LANG", "VERSION", "REPO", "LAST UPDATE"}
	rows := [][]string{
		{"1", "Dagger", "Go", "1.1.8", "https://github.com/mikkurogue/dagger-cli", "Today"},
		{"2", "Black Powder", "Go", "0.0.0", "https://github.com/mikkurogue/black-powder", "Tomorrow"},
		{"3", "Some random Frontend project", "TypeScript", "0.0.0", "https://foo.bar/x", "Never"},
	}

	t := table.New().
		Headers(headers...).
		Rows(rows...).
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}

			if rows[row-1][1] == "Pikachu" {
				return selectedStyle
			}

			switch col {
			case 2, 3: // Type 1 + 2
				c := typeColors
				color := c[fmt.Sprint(rows[row-1][col])]
				return baseStyle.Foreground(color)
			}
			return baseStyle.Foreground(lipgloss.Color("252"))
		}).
		Border(lipgloss.ThickBorder())

	m := Model{t}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
