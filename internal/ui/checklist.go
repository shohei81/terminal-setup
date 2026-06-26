package ui

import (
	"errors"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ErrAborted is returned when the user quits a checklist without confirming.
var ErrAborted = errors.New("user aborted")

type ChecklistItem struct {
	Key      string
	Desc     string
	Selected bool
}

type checklistModel struct {
	title   string
	desc    string
	items   []ChecklistItem
	cursor  int // 0..len(items); len(items) = Confirm button
	aborted bool
}

func (m checklistModel) Init() tea.Cmd { return nil }

func (m checklistModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.aborted = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items) {
				m.cursor++
			}
		case "enter":
			if m.cursor < len(m.items) {
				m.items[m.cursor].Selected = !m.items[m.cursor].Selected
			} else {
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

var (
	clTitleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7aa2f7"))
	clDescStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#565f89"))
	clActiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#c0caf5"))
	clMutedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#565f89"))
	clCheckStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#9ece6a"))
	// Single-line styles (no border) to avoid multi-line prefix alignment issues.
	clBtnActive   = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1a1b26")).
			Background(lipgloss.Color("#7aa2f7")).
			Padding(0, 1)
	clBtnInactive = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3b4261"))
)

func (m checklistModel) View() string {
	var sb strings.Builder

	// Prefix layout (each component is fixed-width):
	//   indent(2) + cursor(2) + check(2) + name(14) + desc
	// Confirm button uses the same indent+cursor+2 spaces to align with name column.
	const indent = "  "

	sb.WriteString("\n")
	sb.WriteString(clTitleStyle.Render(indent+m.title) + "\n")
	sb.WriteString(clDescStyle.Render(indent+m.desc) + "\n\n")

	for i, item := range m.items {
		cur := "  " // 2 chars
		if m.cursor == i {
			cur = "▶ "
		}

		chk := "  " // 2 chars
		if item.Selected {
			chk = clCheckStyle.Render("✓ ")
		}

		name := fmt.Sprintf("%-14s", item.Key)

		if m.cursor == i {
			sb.WriteString(indent + cur + chk + clActiveStyle.Render(name+"  "+item.Desc) + "\n")
		} else {
			sb.WriteString(indent + cur + chk + name + "  " + clMutedStyle.Render(item.Desc) + "\n")
		}
	}

	sb.WriteString("\n")

	// Confirm button — same indent(2)+cursor(2)+pad(2) prefix as item rows.
	// Both states render as a single line so string-prefix alignment holds.
	if m.cursor == len(m.items) {
		sb.WriteString(indent + "▶ " + "  " + clBtnActive.Render(" Confirm ") + "\n")
	} else {
		sb.WriteString(indent + "  " + "  " + clBtnInactive.Render("[ Confirm ]") + "\n")
	}

	sb.WriteString("\n")
	sb.WriteString(clDescStyle.Render(indent+"↑↓ navigate · Enter select/confirm · esc quit") + "\n")

	return sb.String()
}

// RunChecklist runs a bubbletea checklist where Enter toggles items and
// navigating to the Confirm button + Enter submits the selection.
func RunChecklist(title, desc string, items []ChecklistItem) ([]ChecklistItem, error) {
	m := checklistModel{title: title, desc: desc, items: items}
	p := tea.NewProgram(m, tea.WithAltScreen())
	result, err := p.Run()
	if err != nil {
		return nil, err
	}
	final := result.(checklistModel)
	if final.aborted {
		return nil, ErrAborted
	}
	return final.items, nil
}
