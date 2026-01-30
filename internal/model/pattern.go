package model

type Pattern struct {
	Doors     [12]int `json:"doors"`     // index-based internally
	Frequency int     `json:"frequency"` // weight
}
