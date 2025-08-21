package spinner

import (
	"github.com/jgttech/repo/tui/spinner/tui"
)

func SetPending() {
	if !isActive() {
		return
	}

	state.program.Send(
		tui.TuiIcon(""),
	)
}
