package domain

import (
	domain "big2/domain"
	cardPattern "big2/domain/cardPattern"
)

type StraightHandler struct {
	*domain.PlayCardsHandlerBase
}

func NewStraightHandler(next domain.IPlayCardsHandlerBase) *StraightHandler {
	return &StraightHandler{domain.NewPlayCardsHandlerBase(next)}
}

func (p *StraightHandler) ValidatedCardPattern(cards []*domain.Card) (domain.CardPattern, bool){
	straight := cardPattern.NewStraight(cards)
	_, isValidated:= cardPattern.ValidateStraightCard(straight)
	if isValidated {
		return straight, true
	}
	return nil, false
}