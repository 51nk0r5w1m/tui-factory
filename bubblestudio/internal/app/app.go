// Package app is the root Bubble Tea model for the tui-factory shell.
package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuistudio/bubblestudio/internal/keymap"
	"github.com/tuistudio/bubblestudio/internal/list"
	"github.com/tuistudio/bubblestudio/internal/theme"
)

// Model is the root application model.
type Model struct {
	keys     keymap.KeyMap
	theme    theme.Theme
	help     help.Model
	list     list.Model
	width    int
	height   int
	showHelp bool
}

// New returns a ready-to-use Model.
func New() Model {
	return Model{
		keys:  keymap.Default(),
		theme: theme.Default(),
		help:  help.New(),
		list:  list.New(0, 0), // sized on first WindowSizeMsg
	}
}

// Init satisfies tea.Model; no I/O commands on startup.
func (m Model) Init() tea.Cmd {
	return nil
}

// bodyHeight returns the number of lines available for the body region.
// It reserves 1 line each for the header and footer.
func (m Model) bodyHeight() int {
	h := m.height - 2
	if h < 0 {
		h = 0
	}
	return h
}

// Update handles incoming messages and updates state.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(msg.Width, m.bodyHeight())

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.showHelp = !m.showHelp
			return m, nil
		}
	}

	// Forward all messages to the list so j/k/↑/↓ navigation and
	// any internal list updates (including window resizes) are handled.
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the header / body / footer layout.
func (m Model) View() string {
	header := m.theme.Header.Width(m.width).Render("tui-factory")

	var body string
	if m.showHelp {
		body = m.theme.Body.Width(m.width).Height(m.bodyHeight()).Render(m.help.View(m.keys))
	} else {
		// List manages its own sizing; no outer style wrapper needed.
		body = m.list.View()
	}

	footer := m.theme.Footer.Width(m.width).Render("q quit • ? help")

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}
