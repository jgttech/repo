package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/jgttech/repo/tui/theme"
)

func (m Model) View() (view string) {
	if m.Err != nil {
		return m.Err.Error()
	}

	icon := lipgloss.NewStyle().SetString(m.Spinner.View())
	msg := lipgloss.NewStyle().SetString(m.Msg)

	if m.Icon != "" {
		icon = m.IconStyle.SetString(m.Icon)
	}

	if m.Exit && m.Icon == "" {
		icon = m.IconStyle.SetString(theme.DiamondIcon)
	}

	if m.Exit {
		msg = msg.Faint(true)
	}

	view = fmt.Sprintf("%s %s", icon.Render(), msg.Render())

	if m.Exit {
		view += "\n"
		return
	}

	return
}
