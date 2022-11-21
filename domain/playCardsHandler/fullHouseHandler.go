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

func (f *FullHouseHandler) NewCardPattern(cards []*domain.Card) domain.ICardPattern {
	return cardPattern.NewFullHouse(cards)
}
