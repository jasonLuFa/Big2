package domain

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const TotalCardsAmountOfDeck = 52
var topCardIdex = TotalCardsAmountOfDeck - 1

type Deck struct {
	cards [TotalCardsAmountOfDeck]*Card
}

func NewSpecificDeck() *Deck {
	fmt.Println("* <花色> 的值為 C, D, H, S 其中一項，依序代表梅花、菱形、愛心和黑桃。而 <數字> 的值為 3, 4, ..., 9, 10, J, Q, K, A, 2 中其中一項")
	fmt.Println("* 最左邊的牌(<花色>[<數字>]) 表示牌堆最底部的牌，而最右邊的牌為牌堆最上方的牌。在發牌時，會先從最上方的牌開始發牌。")
	fmt.Println("* 以空白隔開每張, 輸入範例 : S[8] S[9] S[3] ... D[5] C[6]")

	cardsSlice := strings.Fields(Scanner("請輸入洗好的牌堆 : "))
	if len(cardsSlice) != 52 {
		panic(fmt.Sprintf("valid number of cards in the deck should be 52, you input %d card",len(cardsSlice)))
	}

	var deck Deck
	for i,cardString := range cardsSlice{
		card := TransformCardStringToCard(cardString)
		deck.cards[i] = card
	}
	return &deck
}

func NewRandomDeck() *Deck {
	ranks := GetRanks()
	suits := GetSuits()

	var deck Deck
	cardAmount := 0
	for _, suit := range suits {
		for _, rank := range ranks {
			card := NewCard(suit,rank)
			deck.cards[cardAmount] = card
			cardAmount++
		}
	}
	deck.shuffle()

	return &deck
}

func (deck *Deck) shuffle() {
	cards := &deck.cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

// func (deck *Deck) GetCard(idx int) *Card {
// 	card := deck.cards[idx]
// 	deck.cards[idx] = nil
// 	return card
// }

func (deck *Deck) Deal() *Card {
	card := deck.cards[topCardIdex]
	deck.cards[topCardIdex] = nil
	topCardIdex--
	return card
}