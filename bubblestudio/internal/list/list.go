// Package list provides a selectable list component for the tui-factory shell.
// Navigation: ↑/↓ or k/j to move, enter to select.
package list

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Item is a single entry shown in the list.
type Item struct {
	title string
	desc  string
}

// NewItem constructs an Item with the given title and description.
func NewItem(title, desc string) Item { return Item{title: title, desc: desc} }

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

// DemoItems is the sample dataset shown in the body.
var DemoItems = []list.Item{
	Item{"Table", "Scrollable data grid with column headers"},
	Item{"Form", "Input fields with validation"},
	Item{"Modal", "Overlay dialog for confirmations"},
	Item{"Progress", "Progress bar or spinner"},
	Item{"Logs", "Streaming log output panel"},
	Item{"Wizard", "Multi-step guided workflow"},
	Item{"Dashboard", "Overview with key metrics"},
	Item{"Status", "Status bar with contextual info"},
}

// Model wraps bubbles/list for use in the app body.
type Model struct {
	list list.Model
}

// New returns a Model sized to w×h with the demo dataset loaded.
func New(w, h int) Model {
	l := list.New(DemoItems, list.NewDefaultDelegate(), w, h)
	l.Title = "Components"
	l.SetShowHelp(false) // help is managed by the parent app
	return Model{list: l}
}

// SetSize resizes the list to fit the available area.
func (m *Model) SetSize(w, h int) {
	m.list.SetSize(w, h)
}

// Selected returns the currently highlighted item, or nil if the list is empty.
func (m Model) Selected() list.Item {
	return m.list.SelectedItem()
}

// Update forwards messages to the inner list model.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the list.
func (m Model) View() string {
	return m.list.View()
}
