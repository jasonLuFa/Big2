package domain

import (
	"big2/domain"
	"fmt"
	"sort"
)

type Straight struct {
	*domain.CardPatternBase
}

func NewStraight(cards []*domain.Card) *Straight {
	cardBase := *domain.NewCardPatternBase("順子", cards)
	return &Straight{&cardBase}
}

func (s *Straight) IsBigger(cardPattern domain.ICardPattern) (bool, error) {
	targetSortedStraightCards, isStraightCard := s.Validate(cardPattern)
	if !isStraightCard {
		return false, fmt.Errorf("your card pattern is not a straight card, so can't compare")
	}

	sort.Slice(s.GetCards(), func(i, j int) bool {
		return !s.GetCards()[i].GetRank().IsBigger(s.GetCards()[j].GetRank())
	})

	return s.GetCards()[4].IsBigger(targetSortedStraightCards.([]*domain.Card)[4]), nil
}

func (s *Straight) Validate(cardPattern domain.ICardPattern) (interface{}, bool) {
	cards := cardPattern.GetCards()
	if len(cards) != 5 {
		return nil, false
	}

	sort.Slice(cards, func(i, j int) bool {
		return !cards[i].GetRank().IsBigger(cards[j].GetRank())
	})

	var isSequential bool
	for i := 0; i <= len(cards)-2; i++ {
		isSequential = cards[i].GetRank().GetRankKind()+1 == cards[i+1].GetRank().GetRankKind()
		// 3(rankKind = 0) Q(rankKind = 9) K A 2
		if cards[i].GetRank().GetRankKind() == 0 && cards[i+1].GetRank().GetRankKind() == 9 {
			isSequential = true
			continue
		}
		// 3 4(rankKind = 1) K(rankKind = 10) A 2
		if cards[i].GetRank().GetRankKind() == 1 && cards[i+1].GetRank().GetRankKind() == 10 {
			isSequential = true
			continue
		}
		// 3 4 5(rankKind = 2) A(rankKind = 11) 2
		if cards[i].GetRank().GetRankKind() == 2 && cards[i+1].GetRank().GetRankKind() == 11 {
			isSequential = true
			continue
		}
		// 3 4 5 6(rankKind = 3) 2(rankKind = 12)
		if cards[i].GetRank().GetRankKind() == 3 && cards[i+1].GetRank().GetRankKind() == 12 {
			isSequential = true
			continue
		}
	}

	return cards, isSequential
}
