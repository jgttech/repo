package card

import "github.com/jgttech/repo/tui/theme"

func Pending(msg string) {
	printf(applyStyle("PENDING", msg, theme.Pending))
}
