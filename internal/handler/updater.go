package handler

import (
	"encoding/json"
	"net/http"

	"case.cubi.bankheist/internal/model"
	"case.cubi.bankheist/internal/service"
)

type UpdateHandler struct {
	patternManager *service.PatternManager
}

type UpdateRequest struct {
	Doors []model.Door `json:"doors"`
}

func (u *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateCompleteDoors(req.Doors); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var final [12]int
	for _, d := range req.Doors {
		idx := (d.Row-1)*4 + (d.Col - 1)
		final[idx] = d.Outcome
	}

	if err := u.patternManager.UpdateFrequencyAndSave(final); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "pattern frequency updated",
	})
}
func NewUpateHandler(patternManager *service.PatternManager) *UpdateHandler {
	return &UpdateHandler{patternManager: patternManager}
}
