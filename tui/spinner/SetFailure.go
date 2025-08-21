package spinner

import (
	"github.com/jgttech/repo/tui/spinner/tui"
	"github.com/jgttech/repo/tui/theme"
)

func SetFailure() {
	if !isActive() {
		return
	}

	state.program.Send(
		tui.TuiIcon(
			theme.Failure.
				SetString(theme.CloseIcon).
				Render(),
		),
	)
}
