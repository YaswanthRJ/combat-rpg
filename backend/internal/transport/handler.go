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

	state, template, err := h.service.StartFight(req.CampaignId, req.Enemy)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	view := ToFightView(state, template)
	json.NewEncoder(w).Encode(view)
}

// POST /fight/action
func (h *Handler) PerformAction(w http.ResponseWriter, r *http.Request) {

	var req struct {
		CampaignId string        `json:"campaignId"`
		Action     domain.Action `json:"action"`
	}
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
