package service

import "case.cubi.bankheist/internal/model"

func GenerateAllPatterns() []model.Pattern {
	results := make([]model.Pattern, 0, 27720)

	counts := map[int]int{
		1: 3, // Diamond
		2: 4, // Cash
		3: 5, // Silver
	}

	var doors [12]int
	backtrack(0, &doors, counts, &results)

	return results
}

func backtrack(
	pos int,
	doors *[12]int,
	counts map[int]int,
	results *[]model.Pattern,
) {
	if pos == 12 {
		p := model.Pattern{
			Doors:     *doors, // array copy (SAFE)
			Frequency: 1,
		}
		*results = append(*results, p)
		return
	}

	for outcome, cnt := range counts {
		if cnt == 0 {
			continue
		}
		doors[pos] = outcome
		counts[outcome]--
		backtrack(pos+1, doors, counts, results)
		counts[outcome]++
	}
}
