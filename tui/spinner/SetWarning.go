package spinner

import (
	"github.com/jgttech/repo/tui/spinner/tui"
	"github.com/jgttech/repo/tui/theme"
)

func SetWarning() {
	if !isActive() {
		return
	}

	state.program.Send(
		tui.TuiIcon(
			theme.Warning.
				SetString(theme.ExclamationIcon).
				Render(),
		),
	)
}
