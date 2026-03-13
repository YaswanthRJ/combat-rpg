package transport

import (
	"combat-sim/internal/app"
)

type Handler struct {
	service *app.CampaignService
}

func NewHandler(service *app.CampaignService) *Handler {
	return &Handler{service: service}
}
