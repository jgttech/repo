package card

import (
	"strings"
	"testing"

	"github.com/jgttech/repo/test"
)

func TestFailureCard(t *testing.T) {
	tag := "tui/card/Failure"
	tests := []string{
		"",
		"t",
		"test 3",
	}

	for _, tt := range tests {
		out := test.Capture(func() {
			Failure(tt)
		})

		if !strings.Contains(out, tt) {
			t.Errorf("[%s] Got message '%s'. Expected '%s'.", tag, out, tt)
		}
	}
}
