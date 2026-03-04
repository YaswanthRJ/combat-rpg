package app

import (
	"combat-sim/internal/domain"
	"testing"
)

func TestStartCampaignCreatesCampaign(t *testing.T) {
	store := NewCampaignStore()
	service := NewCampaignService(store)

	id, err := service.StartCampaign("Bandit")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id == "" {
		t.Errorf("expected campaign ID")
	}

	if _, exists := store.campaigns[id]; !exists {
		t.Errorf("campaign not stored")
	}
}

func TestStartFightWithoutCampaignFails(t *testing.T) {
	store := NewCampaignStore()
	service := NewCampaignService(store)

	_, _, err := service.StartFight("fake", "Bandit")
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestCannotStartFightIfActiveExists(t *testing.T) {
	store := NewCampaignStore()
	service := NewCampaignService(store)

	id, _ := service.StartCampaign("Bandit")

	_, _, err := service.StartFight(id, "Bandit")
	if err != nil {
		t.Fatalf("unexpected error")
	}

	_, _, err = service.StartFight(id, "Bandit")
	if err == nil {
		t.Errorf("expected error when fight already active")
	}
}

func TestPerformActionWithoutFightFails(t *testing.T) {
	store := NewCampaignStore()
	service := NewCampaignService(store)

	id, _ := service.StartCampaign("Bandit")

	_, _, err := service.PerformAction(id, domain.HeavyAttack)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestPlayerHPPersistsAfterFight(t *testing.T) {
	store := NewCampaignStore()
	service := NewCampaignService(store)

	id, _ := service.StartCampaign("Soldier")

	service.StartFight(id, "Bandit")

	// run rounds until fight ends
	for {
		result, _, _ := service.PerformAction(id, domain.HeavyAttack)
		if result.FightEnded {
			break
		}
	}

	campaign := store.campaigns[id]
	if campaign.Player.HP <= 0 {
		t.Errorf("player should survive and persist HP")
	}
}
