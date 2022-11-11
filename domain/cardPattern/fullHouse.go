package domain

import (
	"big2/domain"
	"fmt"
)

type FullHouse struct {
	*domain.CardPatternBase
}

func NewFullHouse(cards []*domain.Card) *FullHouse{
	cardBase := *domain.NewCardPatternBase("葫蘆",cards)
	return &FullHouse{&cardBase}
}

func (f *FullHouse) IsBigger(cardPattern domain.CardPattern) (bool, error) {
	targetFullHouseMap, isPairCard := ValidateFullHouseCard(cardPattern)
	if !isPairCard {
		return false, fmt.Errorf("your card pattern is not a fullHouse card, so can't compare")
	}

	fullHouse := getFullHouseMap(f.GetCards())
	fullHouseCompareRank := getCompareRank(fullHouse)

	targetFullHouseCompareRank := getCompareRank(targetFullHouseMap)

	return fullHouseCompareRank.IsBigger(targetFullHouseCompareRank), nil
}

func ValidateFullHouseCard(cardPattern domain.CardPattern) (map[domain.Rank]int, bool) {
	cards := cardPattern.GetCards()
	if len(cards) != 5{
		return nil,false
	}

	fullHouseMap := getFullHouseMap(cards)

	for _, account := range fullHouseMap {
		if account != 2 && account != 3{
			return nil, false
		}
	}
	return fullHouseMap,true
}

func getFullHouseMap(cards []*domain.Card) map[domain.Rank]int{
	fullHouseMap := make(map[domain.Rank]int)
	for _, card := range cards {
		if _ ,ok:= fullHouseMap[card.GetRank()]; ok{
			fullHouseMap[card.GetRank()] += 1
			continue
		}
		fullHouseMap[card.GetRank()] = 1
	}
	return fullHouseMap
}

func getCompareRank(fullHouseMap map[domain.Rank]int) domain.Rank{
	var fullHouseCompareRank domain.Rank
	for rank, amount := range fullHouseMap{
		if amount ==3 {
			fullHouseCompareRank = rank
		}
	}
	return fullHouseCompareRank
}

