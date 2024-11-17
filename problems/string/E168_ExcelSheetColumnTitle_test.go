package string

import (
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		columnNumber int
		expected     string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
		{2147483647, "FXSHRXW"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := convertToTitle(tt.columnNumber)
			if result != tt.expected {
				t.Errorf("convertToTitle(%d) = %s; want %s", tt.columnNumber, result, tt.expected)
			}
		})
	}
}
