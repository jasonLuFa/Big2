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


func (p *PairHandler) ValidatedCardPattern(cards []*domain.Card) (domain.CardPattern, bool){
	pair := cardPattern.NewPair(cards)
	_, isValidated:= cardPattern.ValidatePairCard(pair)
	if isValidated {
		return pair, true
	}
	return nil, false
}