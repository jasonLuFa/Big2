package domain

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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


func RemoveCardByIndex(cards []*Card, idex int) []*Card {
	newCard := make([]*Card, 0)
	newCard = append(newCard, cards[:idex]...)
	return append(newCard, cards[idex+1:]...)
}
