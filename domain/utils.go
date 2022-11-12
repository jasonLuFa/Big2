package domain

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func Scanner(input string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(input + " ")
	scanner.Scan()
	return scanner.Text()
}

// Transform a string card into Card struct ex : S[8] is a string, it will transform to
// Card{Split, eight}
func TransformCardStringToCard(cardString string) *Card{
	pattern := regexp.MustCompile(`[\]\[]`)
	suitAndSpadeShortcut:= pattern.Split(cardString, -1)
	suitKind := GetSuitFullName(suitAndSpadeShortcut[0]) // S[8] get S -> spade (suitKind)
	rankKind := GetRankFullName(suitAndSpadeShortcut[1]) // S[8] get 8 -> eight (rankKind)
	return NewCard(NewSuit(suitKind),NewRank(rankKind))
}


func RemoveCardsByIdx(cards []*Card, indexes ...int) []*Card {
	sort.Ints(indexes)
	newCard := cards
	for i, v := range indexes{
		tmpCards := make([]*Card,0)
		tmpCards = append(tmpCards, newCard[:v-i]...)
		tmpCards = append(tmpCards, newCard[v-i+1:]...)
		newCard = tmpCards
	}
	return newCard
}
