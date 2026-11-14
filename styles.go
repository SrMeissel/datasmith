package main

import "github.com/charmbracelet/lipgloss"

var (
    BgColor   = lipgloss.Color("#1e1e1e")
    FgColor   = lipgloss.Color("FFD700")
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
	InputStyle: r.NewStyle().Background(BgColor).Foreground(FgColor).Border(lipgloss.RoundedBorder()).BorderForeground(FgColor).Padding(0, 1),
	ChatBoxStyle: r.NewStyle().Background(BgColor).Foreground(FgColor),
	UserTextStyle: r.NewStyle().Align(lipgloss.Right),
	ResponseTextStyle: r.NewStyle().Align(lipgloss.Left),
    }
}
