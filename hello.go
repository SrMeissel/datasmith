package main

import (
    "fmt"
    "strings"
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "os"
)

import "github.com/charmbracelet/lipgloss"
import "github.com/charmbracelet/bubbles/viewport"
import "github.com/charmbracelet/bubbles/textinput"
import "github.com/charmbracelet/bubbletea"

type Message struct {
    Role string `json:"role"`
    Content string `json:"content"`
}

//====================================custom bubbletea messages===============

type streamStartMsg struct {
    reader io.ReadCloser
    decoder *json.Decoder
    err error
}
type streamNextMsg struct {
    content string
    done bool
    reader io.ReadCloser
    decoder *json.Decoder
}

//============================================================================

type Model struct {
    messages []Message

    input textinput.Model
    chatBox viewport.Model

    width int 
    height int
}

func initialModel() Model {
    ti := textinput.New()
    ti.Placeholder = "Start Typing..."
    ti.Focus()
    ti.CharLimit = 500
    ti.Width = 60

    vp := viewport.New(0, 0)
    vp.Style = ChatBoxStyle

    return Model{input: ti, chatBox: vp}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    switch msg := msg.(type) {

    case tea.KeyMsg:
	switch msg.Type {
	
	case tea.KeyEnter:
	    if m.input.Value() != "" {
		m.messages = append(m.messages, Message{Role: "user", Content: m.input.Value() })
		m.input.Reset()

		m.chatBox.SetContent(constructChat(m.messages, m.width))
		m.chatBox.GotoBottom()
		//cmds = append(cmds, fetchAIResponse(m.messages))
		cmds = append(cmds, startOllamaStream(m.messages))
	    }	

	case tea.KeyCtrlC, tea.KeyEsc:
	    return m, tea.Quit
	}

    case tea.WindowSizeMsg:
	m.width = msg.Width
	m.height = msg.Height
	m.input.Width = msg.Width - 4

	m.chatBox.Width = msg.Width
	inputHeight := lipgloss.Height(InputStyle.Render(m.input.View()))
	m.chatBox.Height = msg.Height - inputHeight// - 1

    case streamStartMsg:
	if msg.err != nil {
	    errorText := fmt.Sprintf("Error connecting to ollama: %v", msg.err)
	    m.messages = append(m.messages, Message{Role: "assistant", Content: errorText})
	    m.chatBox.SetContent(constructChat(m.messages, m.width))
	    break
	}
	m.messages = append(m.messages, Message{Role:"assistant", Content: ""})
	cmds = append(cmds, readNextChunk(msg.reader, msg.decoder))
    case streamNextMsg:
	if msg.done{
	    break
	}

	lastIdx := len(m.messages) -1
	m.messages[lastIdx].Content += msg.content

	m.chatBox.SetContent(constructChat(m.messages, m.width))
	m.chatBox.GotoBottom()

	cmds = append(cmds, readNextChunk(msg.reader, msg.decoder))
    }

    //===========================================bubble update==========================================
    var ipcmd, vpcmd tea.Cmd
    m.input, ipcmd = m.input.Update(msg)
    m.chatBox, vpcmd = m.chatBox.Update(msg)
    cmds = append(cmds,ipcmd, vpcmd)

    return m, tea.Batch(cmds...)
}

func (m Model) View() string {
    if m.width == 0 || m.height == 0 {
	return "wait."
    }

    inputView := InputStyle.Render(m.input.View())

    return AppStyle.Render(lipgloss.JoinVertical(lipgloss.Left, m.chatBox.View(), inputView))
}

func startOllamaStream(messages []Message) tea.Cmd {
    return func() tea.Msg {
	reqBody, _ := json.Marshal(map[string]any{
	    "model": "flint",
	    "messages": messages,
	})

	API := os.Getenv("API")
	if API == "" {
	    API = "http://ollama:11434/api/chat"
	}

	resp, err := http.Post(API, "application/json", bytes.NewReader(reqBody))
	if err != nil {
	    return streamStartMsg{err: err}
	}
	return streamStartMsg{ reader: resp.Body, decoder: json.NewDecoder(resp.Body) }
    }
}

func readNextChunk(reader io.ReadCloser, decoder *json.Decoder) tea.Cmd {
    return func() tea.Msg {
	var chunk struct {
	    Message Message `json:"message"`
	    Done bool `json:"done"`
	}
	err := decoder.Decode(&chunk)
	if err != nil || chunk.Done {
	    reader.Close()
	    return streamNextMsg{done: true}
	}

	return streamNextMsg{
	    content: chunk.Message.Content,
	    reader: reader,
	    decoder: decoder,
	}
    }
}

func constructChat(messages[] Message,width int) string {
    UserTextStyle := UserTextStyle.Width(width)
    ResponseTextStyle := ResponseTextStyle.Width(width)

    var chatBuilder strings.Builder
    for _, msg := range messages {
	if msg.Role == "user" {
	    message := "user:\n" + msg.Content
	    chatBuilder.WriteString(UserTextStyle.Render(message))
	    chatBuilder.WriteString("\n\n")
	} else {
	    message := "AI:\n" + msg.Content
	    chatBuilder.WriteString(ResponseTextStyle.Render(message))
	    chatBuilder.WriteString("\n\n")
	}
    }
    return chatBuilder.String()
}

func main() {
    m := initialModel()
    p := tea.NewProgram(m, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
	fmt.Println("Error: ", err);
    }
}
