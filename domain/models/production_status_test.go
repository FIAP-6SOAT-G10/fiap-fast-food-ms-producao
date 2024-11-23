package models

import "testing"

func TestStatusToString(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{Pending, "Pending"},
		{InProgress, "InProgress"},
		{Completed, "Completed"},
		{Failed, "Failed"},
	}

	for _, tt := range tests {
		result := tt.status.String()
		if result != tt.expected {
			t.Errorf("Status.String() = %v, expected %v", result, tt.expected)
		}
	}

}

func TestStatusFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected Status
		wantErr  bool
	}{
		{"Pending", Pending, false},
		{"InProgress", InProgress, false},
		{"Completed", Completed, false},
		{"Failed", Failed, false},
		{"Unknown", 0, true}, // Invalid input case
	}

	for _, tt := range tests {
		result, err := StatusFromString(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("StatusFromString(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
		}
		if err == nil && result != tt.expected {
			t.Errorf("StatusFromString(%q) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}

func TestStatusFromInt(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "Pending"},
		{1, "InProgress"},
		{2, "Completed"},
		{3, "Failed"},
		{4, ""}, // Out-of-range index case
	}

	for _, tt := range tests {
		var result string
		defer func() {
			// Catch panic for invalid index access
			if r := recover(); r != nil {
				if tt.expected != "" {
					t.Errorf("StatusFromInt(%d) panicked unexpectedly", tt.input)
				}
			}
		}()
		result = StatusFromInt(tt.input)
		if result != tt.expected {
			t.Errorf("StatusFromInt(%d) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}
