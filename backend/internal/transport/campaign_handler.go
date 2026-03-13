package transport

import (
	"encoding/json"
	"net/http"
)

// POST /campaign/start
func (h *Handler) StartCampaign(w http.ResponseWriter, r *http.Request) {
	var req StartCampaignRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	id, err := h.service.StartCampaign(req.Creature)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"campaignID": id,
	})

}
