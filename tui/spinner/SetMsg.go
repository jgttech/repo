package spinner

import "github.com/jgttech/repo/tui/spinner/tui"

func SetMsg(msg string) {
	if isActive() {
		state.program.Send(tui.TuiMsg(msg))
	}
}
