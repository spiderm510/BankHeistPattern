// internal/handler/predict.go
package handler

import (
	"encoding/json"
	"net/http"

	"case.cubi.bankheist/internal/model"
	"case.cubi.bankheist/internal/service"
)

type PredictRequest struct {
	Doors []model.Door `json:"doors"`
}

type PredictResponse struct {
	Recommendation model.Door   `json:"recommendation"`
	Doors          []model.Door `json:"doors"`
}

type PredictHandler struct {
	patternManager *service.PatternManager
}

func index(row, col int) int {
	return (row-1)*4 + (col - 1)
}

func (h *PredictHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req PredictRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateDoors(req.Doors); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	revealed := make(map[int]int)
	for _, d := range req.Doors {
		revealed[index(d.Row, d.Col)] = d.Outcome
	}

	best, doors := h.patternManager.Predict(revealed)

	resp := PredictResponse{
		Recommendation: best,
		Doors:          doors,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func NewPredictHandler(patternManager *service.PatternManager) *PredictHandler {
	return &PredictHandler{patternManager: patternManager}
}
