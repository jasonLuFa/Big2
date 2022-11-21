package domain

import (
	domain "big2/domain"
	cardPattern "big2/domain/cardPattern"
)

type PairHandler struct {
	*domain.PlayCardsHandlerBase
}

func NewPairHandler(next domain.IPlayCardsHandlerBase) *PairHandler {
	return &PairHandler{domain.NewPlayCardsHandlerBase(next)}
}

func (p *PairHandler) NewCardPattern(cards []*domain.Card) domain.ICardPattern {
	return cardPattern.NewPair(cards)
}
