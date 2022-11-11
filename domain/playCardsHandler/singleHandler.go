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

func (s *SingleHandler) ValidatedCardPattern(cards []*domain.Card) (domain.CardPattern, bool){
	single := cardPattern.NewSingle(cards)
	_, isValidated := cardPattern.ValidateSingleCard(single)
	if isValidated{
		return single, true
	}
	return nil, false
}