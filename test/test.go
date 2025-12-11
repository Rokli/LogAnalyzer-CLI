package test

import (
	"testing"

	"github.com/Rokli/LogAnalyzer-CLI/internal/types"
)

func TestParserLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected types.LogEntry
		hasError bool
	}{
		{
			name:  "valid log entry",
			input: "2024-01-01 10:00:00 INFO User logged in",
			expected: types.LogEntry{
				Timestamp: "2024-01-01 10:00:00",
				Level:     "INFO",
				Message:   "User logged in",
			},
			hasError: false,
		},
		{
			name:  "valid log entry",
			input: "2023-10-20 10:00:00 ERROR DaTTAtat BaseCONNectred",
			expected: types.LogEntry{
				Timestamp: "2023-10-20 10:00:00",
				Level:     "ERROR",
				Message:   "DaTTAtat BaseCONNectred",
			},
			hasError: false,
		},
		{
			name:  "valid log entry",
			input: "2019-12-22 14:00213:4450 WARN USER protected In b A s E",
			expected: types.LogEntry{
				Timestamp: "2019-12-22 14:00213:4450",
				Level:     "WARN",
				Message:   "USER protected In b A s E",
			},
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := types.ParseLine(tt.input)
			if result != tt.expected && !tt.hasError {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
