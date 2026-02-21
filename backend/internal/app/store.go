package app

import "sync"

type CampaignStore struct {
	mu        sync.Mutex
	campaigns map[string]*Campaign
}

func NewCampaignStore() *CampaignStore {
	return &CampaignStore{
		campaigns: make(map[string]*Campaign),
	}
}
