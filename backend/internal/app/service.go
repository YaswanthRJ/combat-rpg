package app

import (
	"combat-sim/internal/domain"
	"errors"

	"github.com/google/uuid"
)

type FightService struct {
	store *CampaignStore
}

func NewFightService(store *CampaignStore) *FightService {
	return &FightService{store: store}
}

func (s *FightService) StartCampaign(playerCreatureName string) (string, error) {
	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	campaignID := uuid.NewString()

	playerCreature, err := domain.GenerateCreature(playerCreatureName)

	if err != nil {
		return "", err
	}

	campaign := &Campaign{
		ID:     campaignID,
		Player: playerCreature,
		Fights: make(map[string]*domain.FightState),
	}

	s.store.campaigns[campaignID] = campaign

	return campaignID, nil
}

func (s *FightService) StartFight(campaignID string, enemyCreatureName string) (string, error) {
	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	currentCampaign, err := s.getCampaignLocked(campaignID)
	if err != nil {
		return "", err
	}

	if currentCampaign.ActiveFightID != "" {
		activeFight, exists := currentCampaign.Fights[currentCampaign.ActiveFightID]
		if exists && activeFight.FightStatus == domain.Ongoing {
			return "", errors.New("active fight already in progress")
		}
	}

	fightID := uuid.NewString()
	enemyCreature, err := domain.GenerateCreature(enemyCreatureName)

	if err != nil {
		return "", err
	}

	fight := &domain.FightState{
		Player:      currentCampaign.Player,
		Enemy:       enemyCreature,
		FightStatus: domain.Ongoing,
	}

	currentCampaign.Fights[fightID] = fight
	currentCampaign.ActiveFightID = fightID

	return fightID, nil

}

func (s *FightService) PerformAction(campaignID string, action domain.Action) (domain.ActionResult, error) {
	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	currentCampaign, err := s.getCampaignLocked(campaignID)
	if err != nil {
		return domain.ActionResult{}, err
	}

	fight, err := s.getActiveFightLocked(currentCampaign)
	if err != nil {
		return domain.ActionResult{}, err
	}

	result, err := domain.ResolveRound(fight, action)
	if err != nil {
		return domain.ActionResult{}, err
	}

	if fight.FightStatus != domain.Ongoing {
		currentCampaign.Player = fight.Player
	}

	return result, nil
}

func (s *FightService) getCampaignLocked(id string) (*Campaign, error) {
	campaign, ok := s.store.campaigns[id]
	if !ok {
		return nil, errors.New("campaign not found")
	}
	return campaign, nil
}

func (s *FightService) getActiveFightLocked(c *Campaign) (*domain.FightState, error) {

	fightID := c.ActiveFightID
	if fightID == "" {
		return nil, errors.New("no active fight found")
	}

	fight, exists := c.Fights[fightID]
	if !exists {
		return nil, errors.New("active fight missing from campaign")
	}

	if fight.FightStatus != domain.Ongoing {
		return nil, errors.New("no ongoing fight found")
	}

	return fight, nil
}
