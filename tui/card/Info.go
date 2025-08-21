package card

import "github.com/jgttech/repo/tui/theme"

func Info(msg string) {
	printf(applyStyle("INFO", msg, theme.Info))
}
