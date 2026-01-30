package handler

import (
	"errors"
	"fmt"

	"case.cubi.bankheist/internal/model"
)

func validateDoors(doors []model.Door) error {
	if len(doors) == 0 {
		return errors.New("doors array must not be empty")
	}

	seen := make(map[string]bool)

	for _, d := range doors {
		if d.Row < 1 || d.Row > 3 {
			return fmt.Errorf("row must be between 1 and 3 (got %d)", d.Row)
		}
		if d.Col < 1 || d.Col > 4 {
			return fmt.Errorf("col must be between 1 and 4 (got %d)", d.Col)
		}
		if d.Outcome < 1 || d.Outcome > 3 {
			return fmt.Errorf("outcome must be 1, 2, or 3 (got %d)", d.Outcome)
		}

		key := fmt.Sprintf("%d-%d", d.Row, d.Col)
		if seen[key] {
			return fmt.Errorf("duplicate door position (%d,%d)", d.Row, d.Col)
		}
		seen[key] = true
	}

	return nil
}

func validateCompleteDoors(doors []model.Door) error {
	if len(doors) != 12 {
		return errors.New("completed game must contain exactly 12 doors")
	}

	seen := make(map[int]bool)

	for _, d := range doors {
		if d.Row < 1 || d.Row > 3 {
			return fmt.Errorf("row out of range: %d", d.Row)
		}
		if d.Col < 1 || d.Col > 4 {
			return fmt.Errorf("col out of range: %d", d.Col)
		}
		if d.Outcome < 1 || d.Outcome > 3 {
			return fmt.Errorf("invalid outcome: %d", d.Outcome)
		}

		idx := (d.Row-1)*4 + (d.Col - 1)
		if seen[idx] {
			return fmt.Errorf("duplicate door at (%d,%d)", d.Row, d.Col)
		}
		seen[idx] = true
	}
	return nil
}
