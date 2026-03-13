package transport

import (
	"combat-sim/internal/domain"
	"encoding/json"
	"net/http"
)

// POST /fight/action
func (h *Handler) PerformAction(w http.ResponseWriter, r *http.Request) {

	var req PerformActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	if _, ok := domain.ActionPool[req.Action]; !ok {
		http.Error(w, "invalid action", 400)
		return
	}

	result, state, err := h.service.PerformAction(req.CampaignId, req.Action)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	view := ToFightView(state, domain.CreatureTemplate{}) // we no longer have template here

	response := struct {
		Result domain.ActionResult `json:"result"`
		View   FightView           `json:"view"`
	}{
		Result: result,
		View:   view,
	}

	json.NewEncoder(w).Encode(response)
}
