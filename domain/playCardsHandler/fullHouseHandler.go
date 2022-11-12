package domain

import (
	domain "big2/domain"
	cardPattern "big2/domain/cardPattern"
)

type FullHouseHandler struct {
	*domain.PlayCardsHandlerBase
}

func NewFullHouseHandler(next domain.IPlayCardsHandlerBase) *FullHouseHandler {
	return &FullHouseHandler{domain.NewPlayCardsHandlerBase(next)}
}

func (p *FullHouseHandler) ValidatedCardPattern(cards []*domain.Card) (domain.CardPattern, bool){
	fullHouse := cardPattern.NewFullHouse(cards)
	_, isValidated:= cardPattern.ValidateFullHouseCard(fullHouse)
	if isValidated {
		return fullHouse, true
	}
	return nil, false
}