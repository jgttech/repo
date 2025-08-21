package card

import "github.com/jgttech/repo/tui/theme"

func Warning(msg string) {
	printf(applyStyle("WARNING", msg, theme.Warning))
}
