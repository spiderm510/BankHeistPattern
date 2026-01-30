package model

type Door struct {
	Row         int     `json:"row"`
	Col         int     `json:"col"`
	Outcome     int     `json:"outcome"`               // 0 = unknown, 1/2/3 known
	Probability float64 `json:"probability,omitempty"` // used in response
}
