package card

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestApplyStyle(t *testing.T) {
	tests := []struct {
		name  string
		msg   string
		style lipgloss.Style
	}{
		{
			name:  "",
			msg:   "",
			style: lipgloss.NewStyle(),
		},
		{
			name:  "t",
			msg:   "t",
			style: lipgloss.NewStyle(),
		},
		{
			name:  "test",
			msg:   "test",
			style: lipgloss.NewStyle(),
		},
	}

	for _, tt := range tests {
		str := applyStyle(tt.name, tt.msg, tt.style)

		if !strings.Contains(str, tt.name) {
			t.Errorf("Could not find '%s' in name.", tt.name)
		}

		if !strings.Contains(str, tt.msg) {
			t.Errorf("Could not find '%s' in name.", tt.msg)
		}
	}
}
