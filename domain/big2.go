package domain

import (
	"fmt"
	"sort"
)

const numberOfPalyers = 4
var passCount int
var handler PlayCardsHandler

type Big2 struct {
	topPlay   CardPattern
	round     int
	topPlayer *HumanPlayer
	winner    *HumanPlayer
	players   [numberOfPalyers]*HumanPlayer
	deck      *Deck 
}

func NewBig2() *Big2{
	return &Big2{deck: NewSpecificDeck()}
}

func (b *Big2) Start(playCardsHandler PlayCardsHandler) {
	handler = playCardsHandler
	b.initPlayers(playCardsHandler)
	b.dealCardsToPlayers()
	b.takeRound()
}

func (b *Big2) initPlayers(playerCardHandler PlayCardsHandler) {
	for i := 0; i < numberOfPalyers; i++ {
		name := Scanner("What is your name? :")
		players := NewHumanPlayer(name, i, playerCardHandler)
		b.players[i] = players
	}
}

func (b *Big2) dealCardsToPlayers() {
	var playersHandCardsIdexOfClubThree int
	// distribute 4 handCards
	var playersHandCards [4][]*Card
	for cardIdx := 0; cardIdx < HandCardAmount; cardIdx++ {
		for i := range b.players {
			card := b.deck.Deal()
			if card.GetRank().rankKind == 0 && card.GetSuit().suitKind == 0{
				playersHandCardsIdexOfClubThree = i
			}
			playersHandCards[i] = append(playersHandCards[i], card)
		}
	}

	// order handCards
	for _, playerhandCards := range playersHandCards {
		sort.Slice(playerhandCards, func(i, j int) bool {
			return !playerhandCards[i].IsBigger(playerhandCards[j])
		})
	}

	// distribute handCard to each player
	for i, player := range b.players{
		b.players[i].handCards = playersHandCards[i]
		if i == playersHandCardsIdexOfClubThree{
			b.topPlayer = player
		}
	}
}

// CoR 跑過 single 後，之後再跑是從會跳過 single
func (b *Big2) takeRound() {
	isGameOver := false
	for !isGameOver{
		b.round ++
		passCount = 0
		fmt.Printf("新的回合開始了\n\n")

		turnOfPlayerId := b.topPlayer.id
		for passCount < 3 {
			turnOfPlayerId = turnOfPlayerId % 4
			turnOfPlayer := b.players[turnOfPlayerId]
			turnOfPlayer.SetPlayCardsHandler(handler)
			turnOfPlayer.TakeTurn(b.topPlay, b.round)
			turnOfPlayerId++
			
			playerPlayCard := turnOfPlayer.GetPlayCards()
			if playerPlayCard == nil {
				passCount++
				continue
			}

			if len(turnOfPlayer.GetPlayerHandCards()) == 0{
				b.winner = turnOfPlayer
				isGameOver = true
				break
			}
			b.topPlayer = turnOfPlayer
			passCount = 0
			b.topPlay = playerPlayCard
		}
		b.topPlay = nil
	}
}

func (b *Big2) ShowWinner() {
	fmt.Println("遊戲結束，遊戲的勝利者為 ",b.winner.name)
}
