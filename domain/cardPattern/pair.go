package domain

import (
	"big2/domain"
	"fmt"
)

type Pair struct {
	*domain.CardPatternBase
}

func NewPair(cards []*domain.Card) *Pair{
	cardBase := *domain.NewCardPatternBase("對子",cards)
	return &Pair{&cardBase}
}

func (p *Pair) IsBigger(cardPattern domain.CardPattern) (bool, error) {
	targetCards,isPairCard := ValidatePairCard(cardPattern)
	if !isPairCard {
		return false, fmt.Errorf("your card pattern is not a pair card, so can't compare")
	}

	if (p.GetCards()[0].GetRankKind() == targetCards[0].GetRankKind()) {
		return p.sameRankPairIsBiggest()
	}

	return p.GetCards()[0].IsBigger(targetCards[0]), nil
}

func (p *Pair) sameRankPairIsBiggest() (bool, error) {
		if p.GetCards()[0].GetSuitKind() == 3 || p.GetCards()[1].GetSuitKind() == 3 {
			return true, nil
		}
		return false, nil
}

func ValidatePairCard(cardPattern domain.CardPattern) ([]*domain.Card, bool) {
	cards := cardPattern.GetCards()
	if len(cards) == 2 && cards[0].GetRank().IsEqual(cards[1].GetRank()){
		return cards, true
	}
	return nil, false
}