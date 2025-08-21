package card

import (
	"strings"
	"testing"
)

func TestFormatString(t *testing.T) {
	tag := "tui/card/Format"
	tests := [][]string{
		{""},
		{"This", "is", "a", "test"},
	}

	for _, tt := range tests {
		out := Format(tt)

		for _, token := range tt {
			if !strings.Contains(out, token) {
				t.Errorf("[%s] Expected string token '%s', not was not found.", tag, token)
			}
		}
	}
}
