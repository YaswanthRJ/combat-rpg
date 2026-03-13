package transport

import (
	"encoding/json"
	"net/http"
)

// POST /fight/start
func (h *Handler) StartFight(w http.ResponseWriter, r *http.Request) {
	var req StartFightRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	state, template, err := h.service.StartFight(req.CampaignId, req.Enemy)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	view := ToFightView(state, template)
	json.NewEncoder(w).Encode(view)
}
