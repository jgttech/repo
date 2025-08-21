package card

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/jgttech/repo/core/env"
)

func Format(msg []string) string {
	return lipgloss.
		NewStyle().
		SetString(strings.Join(msg, " ")).
		Width(env.TUI_WIDTH).
		Render()
}
