package app

import (
	"combat-sim/internal/domain"
	"errors"

	"github.com/google/uuid"
)

type CampaignService struct {
	store *CampaignStore
}

func NewCampaignService(store *CampaignStore) *CampaignService {
	return &CampaignService{store: store}
}

func (s *CampaignService) StartCampaign(playerCreatureName string) (string, error) {
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

func (s *CampaignService) StartFight(campaignID string, enemyCreatureName string) (*domain.FightState, domain.CreatureTemplate, error) {
	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	currentCampaign, err := s.getCampaignLocked(campaignID)
	if err != nil {
		return nil, domain.CreatureTemplate{}, err
	}

	if currentCampaign.ActiveFightID != "" {
		activeFight, exists := currentCampaign.Fights[currentCampaign.ActiveFightID]
		if exists && activeFight.FightStatus == domain.Ongoing {
			return nil, domain.CreatureTemplate{}, errors.New("active fight already in progress")
		}
	}

	template, ok := domain.CreaturePool[enemyCreatureName]
	if !ok {
		return nil, domain.CreatureTemplate{}, errors.New("unknown enemy type")
	}

	fightID := uuid.NewString()
	enemyCreature, err := domain.GenerateCreature(enemyCreatureName)

	if err != nil {
		return nil, domain.CreatureTemplate{}, err
	}

	fight := &domain.FightState{
		Player:      currentCampaign.Player,
		Enemy:       enemyCreature,
		FightStatus: domain.Ongoing,
	}

	currentCampaign.Fights[fightID] = fight
	currentCampaign.ActiveFightID = fightID

	return fight, template, nil

}

func (s *CampaignService) PerformAction(
	campaignID string,
	action domain.Action,
) (domain.ActionResult, *domain.FightState, error) {

	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	currentCampaign, err := s.getCampaignLocked(campaignID)
	if err != nil {
		return domain.ActionResult{}, nil, err
	}

	fight, err := s.getActiveFightLocked(currentCampaign)
	if err != nil {
		return domain.ActionResult{}, nil, err
	}

	result, err := domain.ResolveRound(fight, action)
	if err != nil {
		return domain.ActionResult{}, nil, err
	}

	if fight.FightStatus != domain.Ongoing {
		currentCampaign.Player = fight.Player
	}

	return result, fight, nil
}

func (s *CampaignService) getCampaignLocked(id string) (*Campaign, error) {
	campaign, ok := s.store.campaigns[id]
	if !ok {
		return nil, errors.New("campaign not found")
	}
	return campaign, nil
}

func (s *CampaignService) getActiveFightLocked(c *Campaign) (*domain.FightState, error) {

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
