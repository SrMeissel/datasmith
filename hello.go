package main

import "fmt"
import "github.com/charmbracelet/lipgloss"
import "github.com/charmbracelet/bubbles/textinput"
import "github.com/charmbracelet/bubbletea"

type Message struct {
    Role string `json:"role"`
    Content string `json:"content"`
}

var (
    bgColor   = lipgloss.Color("#1E1E2E")
    fgColor   = lipgloss.Color("#CDD6F4")
    accent    = lipgloss.Color("#CBA6F7")
    borderCol = lipgloss.Color("63")
)

type Model struct {
    messages []Message
    input textinput.Model
    
    width int 
    height int
}

func initialModel() Model {
    ti := textinput.New()
    ti.Placeholder = "Start Typing..."
    ti.Focus()
    ti.CharLimit = 500
    ti.Width = 60

    return Model{input: ti}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    var cmds []tea.Cmd


    switch msg := msg.(type) {

//==========================================================system commands=================================

    case tea.KeyMsg:
	switch msg.Type {
	
	case tea.KeyEnter:
	    if m.input.Value() != "" {
		m.messages = append(m.messages, Message{Role: "user", Content: m.input.Value() })

		m.input.Reset()

		cmds = append(cmds, fetchAIResponse(m.messages))
	    }	

	case tea.KeyCtrlC, tea.KeyEsc:
	    return m, tea.Quit
	}

//===========================================================================================================

    case tea.WindowSizeMsg:
	m.width = msg.Width
	m.height = msg.Height
	m.input.Width = msg.Width - 4

    case aiResponseMsg:
	if msg.err == nil {
	    m.messages = append(m.messages, Message{Role: "AI", Content: msg.content})
	}
    }

//===========================================textinput bubble update==========================================

    m.input, cmd = m.input.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}

func (m Model) View() string {
    if m.width == 0 || m.height == 0 {
	return "wait."
    }

//======================================compose chat=====================================

    chat := ""
    for _, msg := range m.messages {
	if msg.Role == "user" {
	    chat += "You: " + msg.Content + "\n\n"
	} else {
	    chat += "AI: " + msg.Content + "\n\n"
	}
    }

//========================================styles=========================================

    inputStyle := lipgloss.NewStyle().
    Background(bgColor).
    Border(lipgloss.RoundedBorder()). 
    BorderForeground(borderCol).
    Padding(0, 1).
    Width(m.width - 2)

    inputView := inputStyle.Render(m.input.View())
    chatAreaHeight := m.height - lipgloss.Height(inputView) - 1

    chatBox := lipgloss.NewStyle().
    Width(m.width).
    Height(chatAreaHeight).
    Align(lipgloss.Bottom).
    Render(chat)

    appStyle := lipgloss.NewStyle().
    Background(bgColor).
    Foreground(fgColor).
    Width(m.width).
    Height(m.height)

    return appStyle.Render(lipgloss.JoinVertical(lipgloss.Left, chatBox, inputView))

}

type aiResponseMsg struct {
    content string
    err error
}

func fetchAIResponse(messages []Message) tea.Cmd {
    return func() tea.Msg {
	//call llm endpoint
	content := "hi"
	return aiResponseMsg{content: content}
    }
}

func main() {
    m := initialModel()
    p := tea.NewProgram(m, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
	fmt.Println("Error: ", err);
    }
}
