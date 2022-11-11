package domain

import (
	"big2/domain"
	"fmt"
)

type Single struct {
	*domain.CardPatternBase
}

func NewSingle(cards []*domain.Card) *Single{
	cardBase := *domain.NewCardPatternBase("單張",cards)
	return &Single{&cardBase}
}

func (s *Single) IsBigger(cardPattern domain.CardPattern) (bool, error) {
	targetCard,isSingleCard := ValidateSingleCard(cardPattern)
	if !isSingleCard {
		return false, fmt.Errorf("your card pattern is not a Single card, so can't compare")
	}
	if s.GetCards()[0].IsBigger(targetCard){
		return true, nil
	}
	return false, nil
}

func ValidateSingleCard(cardPattern domain.CardPattern) (*domain.Card,bool) {
	cards := cardPattern.GetCards()
	if len(cards) == 1 {
		return cards[0],true
	}
	return nil,false
}