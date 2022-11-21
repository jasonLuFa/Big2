package domain

import (
	"big2/domain"
	"fmt"
)

type Single struct {
	*domain.CardPatternBase
}

func NewSingle(cards []*domain.Card) *Single {
	cardBase := *domain.NewCardPatternBase("單張", cards)
	return &Single{&cardBase}
}

func (s *Single) IsBigger(cardPattern domain.ICardPattern) (bool, error) {
	targetCard, isSingleCard := s.Validate(cardPattern)
	if !isSingleCard {
		return false, fmt.Errorf("your card pattern is not a Single card, so can't compare")
	}
	if s.GetCards()[0].IsBigger(targetCard.(*domain.Card)) {
		return true, nil
	}
	return false, nil
}

func (s *Single) Validate(cardPattern domain.ICardPattern) (interface{}, bool) {
	cards := cardPattern.GetCards()
	if len(cards) == 1 {
		return cards[0], true
	}
	return nil, false
}
