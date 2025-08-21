package card

import (
	"strings"
	"testing"

	"github.com/jgttech/repo/test"
)

func TestPendingCard(t *testing.T) {
	tag := "tui/card/Pending"
	tests := []string{
		"",
		"t",
		"test 3",
	}

	for _, tt := range tests {
		out := test.Capture(func() {
			Pending(tt)
		})

		if !strings.Contains(out, tt) {
			t.Errorf("[%s] Got message '%s'. Expected '%s'.", tag, out, tt)
		}
	}
}
