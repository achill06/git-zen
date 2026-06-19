package tui

import (
	"fmt"
	"os"

	"github.com/achill06/git-zen/internal/git"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sahilm/fuzzy"
)

var (
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).MarginBottom(1).MarginTop(1).PaddingLeft(2)
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).PaddingLeft(2).Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(lipgloss.Color("212"))
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingLeft(4)
	quitStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginTop(1).PaddingLeft(2)

	// PR Badge Styles
	openPRStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).MarginLeft(2)
	mergedPRStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginLeft(2)
	closedPRStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).MarginLeft(2)
)

type prDataMsg map[string]string

type Model struct {
	choices   []string
	filtered  []string
	cursor    int
	Selected  string
	textInput textinput.Model

	prStatuses map[string]string
	loadingPRs bool
}

func InitialModel() Model {
	branches, err := git.GetRecentBranches()
	if err != nil {
		fmt.Printf("Error fetching branches: %v\n", err)
		os.Exit(1)
	}

	ti := textinput.New()
	ti.Placeholder = "Search branches..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	return Model{
		choices:    branches,
		filtered:   branches,
		cursor:     0,
		textInput:  ti,
		prStatuses: make(map[string]string),
		loadingPRs: true,
	}
}

func fetchPRs() tea.Msg {
	statuses, err := git.GetPRStatuses()
	if err != nil {
		return prDataMsg(make(map[string]string))
	}
	return prDataMsg(statuses)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, fetchPRs)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case prDataMsg:
		m.prStatuses = msg
		m.loadingPRs = false
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyUp, tea.KeyCtrlK:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown, tea.KeyCtrlJ:
			if m.cursor < len(m.filtered)-1 {
				m.cursor++
			}
		case tea.KeyEnter:
			if len(m.filtered) > 0 {
				m.Selected = m.filtered[m.cursor]
			}
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)

	query := m.textInput.Value()
	if query == "" {
		m.filtered = m.choices
	} else {
		matches := fuzzy.Find(query, m.choices)
		m.filtered = make([]string, len(matches))
		for i, match := range matches {
			m.filtered[i] = match.Str
		}
	}

	if m.cursor >= len(m.filtered) {
		m.cursor = max(0, len(m.filtered)-1)
	}

	return m, cmd
}

func (m Model) View() string {
	s := titleStyle.Render("Fuzzy Switch Branch:") + "\n\n"
	s += "  " + m.textInput.View() + "\n\n"

	for i, choice := range m.filtered {
		badge := ""
		if state, exists := m.prStatuses[choice]; exists {
			switch state {
			case "OPEN":
				badge = openPRStyle.Render("● Open")
			case "MERGED":
				badge = mergedPRStyle.Render("● Merged")
			case "CLOSED":
				badge = closedPRStyle.Render("● Closed")
			}
		}

		line := choice + badge

		if m.cursor == i {
			s += selectedStyle.Render(line) + "\n"
		} else {
			s += normalStyle.Render(line) + "\n"
		}
	}

	// Loading indicator for network request
	if m.loadingPRs {
		s += quitStyle.Render("Loading GitHub PR data...") + "\n"
	} else {
		s += quitStyle.Render("Press esc to quit.") + "\n"
	}

	return s
}

func Run() (string, error) {
	p := tea.NewProgram(InitialModel())
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}
	if m, ok := finalModel.(Model); ok && m.Selected != "" {
		return m.Selected, nil
	}
	return "", nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
