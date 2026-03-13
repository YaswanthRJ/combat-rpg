package transport

import "combat-sim/internal/domain"

type StartCampaignRequest struct {
	Creature string `json:"Creature"`
}

type StartFightRequest struct {
	CampaignId string `json:"campaignId"`
	Enemy      string `json:"enemy"`
}

type PerformActionRequest struct {
	CampaignId string        `json:"campaignId"`
	Action     domain.Action `json:"action"`
}
