package main

import "github.com/charmbracelet/lipgloss"

var (
    BgColor   = lipgloss.Color("#111111")
    FgColor   = lipgloss.Color("#FFA000")

    MarkdownTheme = 
	`{
	"document": {
	    "color": "#FFA000",
	    "block_prefix": "\n",
	    "block_suffix": "\n"
	},
	"block_quote": {
	    "color": "#FFA500",
	    "indent": 2,
	    "indent_token": "│ "
	},
	"paragraph": {
	    "color": "#FFA000",
	    "background_color": "#111111"
	},
	"heading": {
	    "color": "#FFFFFF",
	    "background_color": "#FF8C00",
	    "bold": true,
	    "block_suffix": "\n"
	},
	"strong": {
	    "color": "#FF8C00",
	    "bold": true
	},
	"emph": {
	    "color": "#FFA500",
	    "italic": true
	},
	"code": {
	    "color": "#FFFFFF",
	    "background_color": "#8B4500"
	},
	"code_block": {
	    "color": "#FFD700",
	    "background_color": "#1e1e1e",
	    "margin": 2
	},
	"link": {
	    "color": "#FF8C00",
	    "underline": true
	},
	"list": {
	    "level_indent": 2
	},
	"item": {
	    "block_prefix": "• "
	}
	}`
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
