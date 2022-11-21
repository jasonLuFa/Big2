package domain

import (
	"big2/domain"
	"fmt"
)

type FullHouse struct {
	*domain.CardPatternBase
}

func NewFullHouse(cards []*domain.Card) *FullHouse {
	cardBase := *domain.NewCardPatternBase("葫蘆", cards)
	return &FullHouse{&cardBase}
}

func (f *FullHouse) IsBigger(cardPattern domain.ICardPattern) (bool, error) {
	targetFullHouseMap, isPairCard := f.Validate(cardPattern)
	if !isPairCard {
		return false, fmt.Errorf("your card pattern is not a fullHouse card, so can't compare")
	}

	fullHouse := f.getFullHouseMap(f.GetCards())
	fullHouseCompareRank := f.getCompareRank(fullHouse)

	targetFullHouseCompareRank := f.getCompareRank(targetFullHouseMap.(map[domain.Rank]int))

	return fullHouseCompareRank.IsBigger(targetFullHouseCompareRank), nil
}

func (f *FullHouse) Validate(cardPattern domain.ICardPattern) (interface{}, bool) {
	cards := cardPattern.GetCards()
	if len(cards) != 5 {
		return nil, false
	}

	fullHouseMap := f.getFullHouseMap(cards)

	for _, account := range fullHouseMap {
		if account != 2 && account != 3 {
			return nil, false
		}
	}
	return fullHouseMap, true
}

func (f *FullHouse) getFullHouseMap(cards []*domain.Card) map[domain.Rank]int {
	fullHouseMap := make(map[domain.Rank]int)
	for _, card := range cards {
		if _, ok := fullHouseMap[card.GetRank()]; ok {
			fullHouseMap[card.GetRank()] += 1
			continue
		}
		fullHouseMap[card.GetRank()] = 1
	}
	return fullHouseMap
}

func (f *FullHouse) getCompareRank(fullHouseMap map[domain.Rank]int) domain.Rank {
	var fullHouseCompareRank domain.Rank
	for rank, amount := range fullHouseMap {
		if amount == 3 {
			fullHouseCompareRank = rank
		}
	}
	return fullHouseCompareRank
}
