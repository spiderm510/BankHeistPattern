package service

import (
	"log"

	"case.cubi.bankheist/internal/model"
	"case.cubi.bankheist/internal/storage"
)

type PatternManager struct {
	patterns []model.Pattern
	storage  storage.Storage
}

func NewPatternManager(storage storage.Storage) *PatternManager {
	log.Println("Loading patterns...")
	patterns, err := storage.Load()
	if err != nil {
		log.Println("patterns.json not found, generating patterns...")

		patterns = GenerateAllPatterns()

		log.Printf("Generated %d patterns\n", len(patterns))

		if err := storage.Save(patterns); err != nil {
			log.Fatal("Failed to save patterns:", err)
		}

		log.Println("Patterns saved")
	} else {
		log.Println("Patterns loaded")
	}
	return &PatternManager{patterns: patterns, storage: storage}
}

func (pm *PatternManager) Match(revealed map[int]int) []model.Pattern {
	var matched []model.Pattern
	for _, p := range pm.patterns {
		ok := true
		for idx, outcome := range revealed {
			if p.Doors[idx] != outcome {
				ok = false
				break
			}
		}
		if ok {
			matched = append(matched, p)
		}
	}
	return matched
}

func (pm *PatternManager) Predict(revealed map[int]int) (model.Door, []model.Door) {
	matched := pm.Match(revealed)

	totalWeight := pm.TotalWeight(matched)
	if totalWeight == 0 {
		return model.Door{}, nil
	}

	var (
		bestProb = -1.0
		best     model.Door
		results  []model.Door
	)

	for i := 0; i < 12; i++ {
		if _, revealedAlready := revealed[i]; revealedAlready {
			continue
		}

		prob := pm.DiamondProbabilityAtIndex(matched, totalWeight, i)

		d := model.Door{
			Row:         i/4 + 1,
			Col:         i%4 + 1,
			Outcome:     1, // Diamond
			Probability: prob,
		}

		results = append(results, d)

		if prob > bestProb {
			bestProb = prob
			best = d
		}
	}

	return best, results
}
func (pm *PatternManager) TotalWeight(patterns []model.Pattern) int {
	total := 0
	for _, p := range patterns {
		total += p.Frequency
	}
	return total
}
func (pm *PatternManager) DiamondProbabilityAtIndex(
	patterns []model.Pattern,
	totalWeight int,
	index int,
) float64 {

	if totalWeight == 0 {
		return 0
	}

	diamondWeight := 0
	for _, p := range patterns {
		if p.Doors[index] == 1 {
			diamondWeight += p.Frequency
		}
	}

	return float64(diamondWeight) / float64(totalWeight)
}

func (pm *PatternManager) UpdateFrequency(final [12]int) {
	for i := range pm.patterns {
		if pm.patterns[i].Doors == final {
			pm.patterns[i].Frequency++
			return
		}
	}
}
func (pm *PatternManager) UpdateFrequencyAndSave(final [12]int) error {
	pm.UpdateFrequency(final)
	return pm.storage.Save(pm.patterns)
}

func (pm *PatternManager) Patterns() []model.Pattern {
	return pm.patterns
}
