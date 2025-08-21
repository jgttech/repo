package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

type TuiOption func(*Model)
type TuiError error
type TuiMsg string
type TuiIcon string

type Model struct {
	Spinner      spinner.Model
	PendingStyle lipgloss.Style
	IconStyle    lipgloss.Style
	Icon         string
	Msg          string
	Exit         bool
	Err          error
}
