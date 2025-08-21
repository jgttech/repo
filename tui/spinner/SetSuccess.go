package spinner

import (
	"github.com/jgttech/repo/tui/spinner/tui"
	"github.com/jgttech/repo/tui/theme"
)

func SetSuccess() {
	if !isActive() {
		return
	}

	state.program.Send(
		tui.TuiIcon(
			theme.Success.
				SetString(theme.CheckIcon).
				Render(),
		),
	)
}
