package theme

import "github.com/charmbracelet/lipgloss"

const (
	// Icons
	CheckIcon       = " ✓ "
	CloseIcon       = " 𐄂 "
	DiamondIcon     = " ❖ "
	ExclamationIcon = " ! "

	// Theme
	AmberYellow = lipgloss.Color("#ffa726")
	DarkAmber   = lipgloss.Color("#4a3d1f")
	SkyBlue     = lipgloss.Color("#5fb3f0")
	DeepBlue    = lipgloss.Color("#1f2f4a")
	MintGreen   = lipgloss.Color("#4ade80")
	DarkForest  = lipgloss.Color("#1f4a2f")
	CrimsonRed  = lipgloss.Color("#dc2626")
	TealBlue    = lipgloss.Color("#14b8a6")
)

var (
	Info    = lipgloss.NewStyle().Foreground(SkyBlue).Background(DeepBlue)
	Pending = lipgloss.NewStyle().Foreground(TealBlue)
	Success = lipgloss.NewStyle().Foreground(MintGreen).Background(DarkForest)
	Warning = lipgloss.NewStyle().Foreground(AmberYellow).Background(DarkAmber)
	Failure = lipgloss.NewStyle().Foreground(CrimsonRed).Background(nil)
)
