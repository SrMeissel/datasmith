package main

import "github.com/charmbracelet/lipgloss"

var (
    BgColor   = lipgloss.Color("#1E1E2E")
    FgColor   = lipgloss.Color("#CDD6F4")
    Accent    = lipgloss.Color("#CBA6F7")
    BorderCol = lipgloss.Color("63")

    InputStyle = lipgloss.NewStyle().
    Background(BgColor).
    Border(lipgloss.RoundedBorder()). 
    BorderForeground(BorderCol).
    Padding(0, 1)
    //Width(m.width - 2)

    ChatBoxStyle = lipgloss.NewStyle()
    //Align(lipgloss.Left)
    //Width(m.width).
    //Height(chatAreaHeight).
    //Render(chat)
    
    ResponseTextStyle = lipgloss.NewStyle().
    Align(lipgloss.Left)

    UserTextStyle = lipgloss.NewStyle().
    Align(lipgloss.Right)

    AppStyle = lipgloss.NewStyle().
    Background(BgColor).
    Foreground(FgColor)
    //Width(m.width).
    //Height(m.height)
)
