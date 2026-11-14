package main

import "github.com/charmbracelet/lipgloss"

var (
    BgColor   = lipgloss.Color("#111111")
    FgColor   = lipgloss.Color("#FFA000")
)

type Styles struct {
    AppStyle lipgloss.Style

    InputStyle lipgloss.Style
    ChatBoxStyle lipgloss.Style

    UserTextStyle lipgloss.Style
    ResponseTextStyle lipgloss.Style
}

func BuildStyles(r *lipgloss.Renderer) Styles {
    return Styles{
	AppStyle: r.NewStyle(),
	InputStyle: r.NewStyle().Background(BgColor).Foreground(FgColor).Border(lipgloss.RoundedBorder()).BorderBackground(BgColor).BorderForeground(FgColor).Padding(0, 1),
	ChatBoxStyle: r.NewStyle().Background(BgColor).Foreground(FgColor),
	UserTextStyle: r.NewStyle().Align(lipgloss.Right),
	ResponseTextStyle: r.NewStyle().Align(lipgloss.Left),
    }
}
