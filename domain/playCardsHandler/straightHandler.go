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

func (s *StraightHandler) NewCardPattern(cards []*domain.Card) domain.ICardPattern {
	return cardPattern.NewStraight(cards)
}
