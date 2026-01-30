package service

import (
	"testing"

	"case.cubi.bankheist/internal/model"
	"github.com/stretchr/testify/require"
)

func TestPatternMatching(t *testing.T) {
	patterns := []model.Pattern{
		{Doors: [12]int{1, 2, 3}, Frequency: 1},
		{Doors: [12]int{1, 3, 3}, Frequency: 1},
		{Doors: [12]int{2, 2, 3}, Frequency: 1},
	}

	revealed := map[int]int{
		0: 1, // (1,1)
		2: 3, // (1,3)
	}
	pm := &PatternManager{
		patterns: patterns,
	}
	matched := pm.Match(revealed)

	if len(matched) != 2 {
		t.Fatalf("expected 2 matched patterns, got %d", len(matched))
	}

	for _, p := range matched {
		if p.Doors[0] != 1 || p.Doors[2] != 3 {
			t.Fatalf("pattern does not match revealed constraints")
		}
	}
}
func TestPredict(t *testing.T) {
	pm := &PatternManager{
		patterns: []model.Pattern{
			{Doors: [12]int{1, 2, 3}, Frequency: 1},
			{Doors: [12]int{2, 2, 3}, Frequency: 1},
		},
	}
	revealed := map[int]int{}

	best, doors := pm.Predict(revealed)
	require.Len(t, doors, 12)
	var door0 model.Door
	for _, d := range doors {
		if d.Row == 1 && d.Col == 1 {
			door0 = d
			break
		}
	}
	require.NotZero(t, door0)
	require.InEpsilon(t, 0.5, door0.Probability, 0.0001)
	require.True(t, best.Probability >= 0.5)
}

func TestUpdateFrequency(t *testing.T) {
	pm := &PatternManager{
		patterns: []model.Pattern{
			{Doors: [12]int{1, 2, 3}, Frequency: 1},
			{Doors: [12]int{2, 2, 3}, Frequency: 1},
		},
	}

	target := [12]int{1, 2, 3}
	pm.UpdateFrequency(target)

	if pm.patterns[0].Frequency != 2 {
		t.Fatalf("expected frequency 2, got %d", pm.patterns[0].Frequency)
	}

	if pm.patterns[1].Frequency != 1 {
		t.Fatalf("unexpected frequency update on non-matching pattern")
	}
}
