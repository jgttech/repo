package spinner

import (
	"github.com/jgttech/repo/tui/spinner/tui"
	"github.com/jgttech/repo/tui/theme"
)

func SetInfo() {
	if !isActive() {
		return
	}

	state.program.Send(
		tui.TuiIcon(
			theme.Info.
				SetString(theme.DiamondIcon).
				Render(),
		),
	)
}
