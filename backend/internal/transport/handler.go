package transport

import (
	"combat-sim/internal/app"
	"combat-sim/internal/domain"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *app.CampaignService
}

func NewHandler(service *app.CampaignService) *Handler {
	return &Handler{service: service}
}

// POST /campaign/start
func (h *Handler) StartCampaign(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Creature string `json:"Creature"`
	}

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

// POST /fight/start
func (h *Handler) StartFight(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CampaignId string `json:"campaignId"`
		Enemy      string `json:"enemy"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	template, ok := domain.CreaturePool[req.Enemy]
	if !ok {
		http.Error(w, "unknown enemy type", 400)
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
