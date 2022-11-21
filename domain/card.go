package domain

import (
	"fmt"
)

type rankKind int
type suitKind int

const (
	three rankKind = iota
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
	two
)

const (
	club    suitKind = iota // 梅花(C)
	diamond                 // 菱形(D)
	heart                   // 愛心(H)
	spade                   // 黑桃(S)
)

// card
type Card struct {
	Suit
	Rank
}

func NewCard(s Suit, r Rank) *Card {
	return &Card{s, r}
}

func (card *Card) String() string {
	return fmt.Sprintf("%s[%s]", card.GetSuit().GetSuitKind().String(), card.GetRank().GetRankKind().String())
}

func (card *Card) GetRank() Rank {
	return card.Rank
}

func (card *Card) GetSuit() Suit {
	return card.Suit
}

func (card *Card) IsBigger(targetCard *Card) bool {
	return (card.GetRank().IsBigger(targetCard.GetRank())) || (card.GetRank().IsEqual(targetCard.GetRank()) && card.GetSuit().IsBigger(targetCard.GetSuit()))
}

// suit
func (sk suitKind) String() string {
	return []string{"C", "D", "H", "S"}[sk]
}

type Suit struct {
	suitKind
}

func NewSuit(sk suitKind) Suit {
	return Suit{sk}
}

func GetSuits() [4]Suit {
	return [4]Suit{
		{club},
		{diamond},
		{heart},
		{spade}}
}

func (s Suit) GetSuitKind() suitKind {
	return s.suitKind
}

func (s Suit) IsBigger(targetSuit Suit) bool {
	return s.GetSuitKind() > targetSuit.GetSuitKind()
}

func GetSuitFullName(suitShortcut string) suitKind {
	switch suitShortcut {
	case "C":
		return club
	case "D":
		return diamond
	case "H":
		return heart
	case "S":
		return spade
	default:
		panic(fmt.Sprintf("there is no such suit shortcut, %s", suitShortcut))
	}
}

// Rank
type Rank struct {
	rankKind
}

func NewRank(rk rankKind) Rank {
	return Rank{rk}
}

func (rk rankKind) String() string {
	return []string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}[rk]
}

func GetRanks() [13]Rank {
	return [13]Rank{
		{two},
		{three},
		{four},
		{five},
		{six},
		{seven},
		{eight},
		{nine},
		{ten},
		{jack},
		{queen},
		{king},
		{ace}}
}

func (r Rank) GetRankKind() rankKind {
	return r.rankKind
}

func (r Rank) IsBigger(targetRank Rank) bool {
	return r.GetRankKind() > targetRank.GetRankKind()
}

func (r Rank) IsEqual(targetRank Rank) bool {
	return r.GetRankKind() == targetRank.GetRankKind()
}

func GetRankFullName(rankShortcut string) rankKind {
	switch rankShortcut {
	case "3":
		return three
	case "4":
		return four
	case "5":
		return five
	case "6":
		return six
	case "7":
		return seven
	case "8":
		return eight
	case "9":
		return nine
	case "10":
		return ten
	case "J":
		return jack
	case "Q":
		return queen
	case "K":
		return king
	case "A":
		return ace
	case "2":
		return two
	default:
		panic(fmt.Sprintf("there is no such rank shortcut, %s", rankShortcut))
	}
}

// cardPattern
type CardPatern struct {
	ICardPattern
}

type ICardPattern interface {
	IsBigger(ICardPattern) (bool, error)
	GetCards() []*Card
	GetKind() string
	Validate(ICardPattern) (interface{}, bool)
}

type CardPatternBase struct {
	kind  string
	cards []*Card
}

func NewCardPatternBase(kind string, cards []*Card) *CardPatternBase {
	return &CardPatternBase{kind: kind, cards: cards}
}

func (c *CardPatternBase) GetCards() []*Card {
	return c.cards
}

func (c *CardPatternBase) GetKind() string {
	return c.kind
}
