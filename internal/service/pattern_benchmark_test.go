package service

import (
	"testing"

	"case.cubi.bankheist/internal/model"
)

func BenchmarkPatternMatching(b *testing.B) {
	patterns := generateTestPatterns(27720)

	pm := &PatternManager{
		patterns: patterns,
	}

	revealed := map[int]int{
		0: 1,
		5: 3,
		7: 2,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = pm.Match(revealed)
	}
}
func BenchmarkDiamondProbability(b *testing.B) {
	patterns := generateTestPatterns(10000)

	pm := &PatternManager{
		patterns: patterns,
	}

	matched := pm.Match(map[int]int{0: 1})

	total := pm.TotalWeight(matched)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = pm.DiamondProbabilityAtIndex(matched, total, 4)
	}
}

func generateTestPatterns(limit int) []model.Pattern {
	base := []int{
		1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3, 3,
	}

	patterns := make([]model.Pattern, 0, limit)

	var permute func(int)
	permute = func(i int) {
		if len(patterns) >= limit {
			return
		}

		if i == len(base) {
			var doors [12]int
			copy(doors[:], base)

			patterns = append(patterns, model.Pattern{
				Doors:     doors,
				Frequency: 1,
			})
			return
		}

		used := make(map[int]bool)
		for j := i; j < len(base); j++ {
			if used[base[j]] {
				continue
			}
			used[base[j]] = true

			base[i], base[j] = base[j], base[i]
			permute(i + 1)
			base[i], base[j] = base[j], base[i]

			if len(patterns) >= limit {
				return
			}
		}
	}

	permute(0)
	return patterns
}
