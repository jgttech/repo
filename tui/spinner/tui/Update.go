package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.Exit = true
			return m, tea.Quit
		default:
			return m, nil
		}
	case TuiError:
		m.Err = msg
		return m, nil
	case TuiMsg:
		m.Msg = string(msg)
		return m, nil
	case TuiIcon:
		m.Icon = string(msg)
		return m, nil
	default:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}
