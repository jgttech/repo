package card

import "github.com/jgttech/repo/tui/theme"

func Failure(msg string) {
	printf(applyStyle("FAILURE", msg, theme.Failure))
}
