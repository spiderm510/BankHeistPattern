package service

import "testing"

func TestPatternInitialization_UniqueAndCount(t *testing.T) {
	patterns := GenerateAllPatterns()

	if len(patterns) != 27720 {
		t.Fatalf("expected 27720 patterns, got %d", len(patterns))
	}

	seen := make(map[[12]int]bool)
	for _, p := range patterns {
		if seen[p.Doors] {
			t.Fatalf("duplicate pattern detected: %+v", p.Doors)
		}
		seen[p.Doors] = true
	}

	for _, p := range patterns {
		if p.Frequency != 1 {
			t.Fatalf("expected initial frequency 1, got %d", p.Frequency)
		}
	}
}
