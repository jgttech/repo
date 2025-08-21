package card

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/jgttech/repo/core/env"
)

func applyStyle(title, msg string, style lipgloss.Style) (view string) {
	timetamp := time.Now().UTC().Format(time.UnixDate)
	header := title + "  " + timetamp

	text := lipgloss.
		NewStyle().
		Width(env.TUI_WIDTH).
		Bold(true).
		BorderStyle(lipgloss.RoundedBorder()).
		PaddingLeft(1).
		PaddingRight(1).
		Render(header) + "\n"

	msg = lipgloss.
		NewStyle().
		Bold(true).
		Render("MESSAGE") + "\n\n" + msg

	text += lipgloss.
		NewStyle().
		Width(env.TUI_WIDTH).
		BorderStyle(lipgloss.RoundedBorder()).
		PaddingLeft(1).
		PaddingRight(1).
		Render(msg)

	view = lipgloss.
		NewStyle().
		PaddingTop(1).
		Render(text)

	return
}
