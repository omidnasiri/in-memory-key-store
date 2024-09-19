package ks

import (
	"testing"
)

func TestSimpleHashFunc(t *testing.T) {
	tests := []struct {
		key      string
		limit    int
		expected int
	}{
		{"test", 16, 0},
		{"hello", 16, 4},
		{"world", 16, 8},
		{"golang", 16, 8},
		{"hash", 16, 4},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			result := SimpleHashFunc(tt.key, tt.limit)
			if result != tt.expected {
				t.Errorf("SimpleHashFunc(%s, %d) = %d; want %d", tt.key, tt.limit, result, tt.expected)
			}
		})
	}
}

func TestFowlerNollVoHashFunction(t *testing.T) {
	tests := []struct {
		key      string
		limit    int
		expected int
	}{
		{"test", 16, 5},
		{"hello", 16, 11},
		{"world", 16, 3},
		{"golang", 16, 3},
		{"hash", 16, 1},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			result := FowlerNollVoHashFunction(tt.key, tt.limit)
			if result != tt.expected {
				t.Errorf("FowlerNollVoHashFunction(%s, %d) = %d; want %d", tt.key, tt.limit, result, tt.expected)
			}
		})
	}
}
