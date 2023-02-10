package app
import ("github.com/charmbracelet/lipgloss")

var interfaceStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("63")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(2).Margin(1)

var servingStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#3079CF")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(2).Margin(1)

var errStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#f55050")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(2).
	Margin(1)

var logStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#64C964")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(0).Margin(0)
