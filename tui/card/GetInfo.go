package card

import (
	"fmt"

	"github.com/jgttech/repo/tui/theme"
)

func GetInfo(msg string, a ...any) string {
	msg = fmt.Sprintf(msg, a...)
	return sprintf(applyStyle("INFO", msg, theme.Info))
}
