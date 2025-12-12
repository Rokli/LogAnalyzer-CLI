package test

import (
	"testing"

	"github.com/Rokli/LogAnalyzer-CLI/internal/types"
	"github.com/Rokli/LogAnalyzer-CLI/pkg/logs"
)

func TestParserLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected logs.LogEntry
		hasError bool
	}{
		{
			name:  "valid log entry",
			input: "2024-01-01 10:00:00 INFO User logged in",
			expected: logs.LogEntry{
				Timestamp: "2024-01-01 10:00:00",
				Level:     "INFO",
				Message:   "User logged in",
			},
			hasError: false,
		},
		{
			name:  "valid log entry",
			input: "2023-10-20 10:00:00 ERROR DaTTAtat BaseCONNectred",
			expected: logs.LogEntry{
				Timestamp: "2023-10-20 10:00:00",
				Level:     "ERROR",
				Message:   "DaTTAtat BaseCONNectred",
			},
			hasError: false,
		},
		{
			name:  "valid log entry",
			input: "2019-12-22 14:00213:4450 WARN USER protected In b A s E",
			expected: logs.LogEntry{
				Timestamp: "2019-12-22 14:00213:4450",
				Level:     "WARN",
				Message:   "USER protected In b A s E",
			},
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := types.ParseLine(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestGetStatsEmpty(t *testing.T) {
	result := types.GetStats([]logs.LogEntry{})
	expected := map[string]int{"INFO": 0, "ERROR": 0, "WARN": 0}

	for level, count := range expected {
		if result[level] != count {
			t.Errorf("expected %s: %d, got %d", level, count, result[level])
		}
	}
}

func TestGetLimitStrLimitExceeds(t *testing.T) {
	logs := []logs.LogEntry{
		{Timestamp: "1", Level: "INFO", Message: "test1"},
		{Timestamp: "2", Level: "INFO", Message: "test2"},
	}

	result := types.GetLimitStr(logs, 5)
	if len(result) != 2 {
		t.Errorf("expected 2 logs when limit exceeds, got %d", len(result))
	}
}
