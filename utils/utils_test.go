package utils

import (
	"testing"
)

func TestNewLogger(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "base",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(); got == nil {
				t.Errorf("NewLogger() should not return a nil logger")
			}
		})
	}
}
