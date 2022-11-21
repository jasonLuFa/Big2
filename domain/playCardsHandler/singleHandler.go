package domain

import (
	domain "big2/domain"
	cardPattern "big2/domain/cardPattern"
)

type SingleHandler struct {
	*domain.PlayCardsHandlerBase
}

func NewSingleHandler(next domain.IPlayCardsHandlerBase) *SingleHandler {
	return &SingleHandler{domain.NewPlayCardsHandlerBase(next)}
}

func (s *SingleHandler) NewCardPattern(cards []*domain.Card) domain.ICardPattern {
	return cardPattern.NewSingle(cards)
}
