package domain

import (
	"fmt"
	"strconv"
	"strings"
)

const HandCardAmount = 13
var Turn = 1

type HumanPlayer struct {
	name      string
	id        int
	handCards []*Card
	playerCardsHandler PlayCardsHandler
	playCards CardPattern
}

func NewHumanPlayer(name string, id int, playerCardsHandler PlayCardsHandler) *HumanPlayer {
	var emptyHandCards []*Card
	humanPlayer := &HumanPlayer{name :name, id:id, handCards:emptyHandCards, playerCardsHandler:playerCardsHandler}
	return humanPlayer
}

func (player *HumanPlayer) TakeTurn(topCardPattern CardPattern, round int) {
	fmt.Printf("輪到%s了\n",player.name)
	fmt.Println(player.name,"的手牌 :", player.handCards)

	var playCardsIdx []int
	isValidated := false
	var playCardPattern CardPattern
	for !isValidated{
		var playCards []*Card
		playCards, playCardsIdx = player.play()
		// palyer pass
		if playCards == nil && playCardsIdx == nil{
			if Turn == 1 {
				fmt.Println("你不能在新的回合中喊 PASS")
				continue
			}
			fmt.Printf("玩家 %s PASS\n\n",player.name)
			player.playCards = nil
			Turn++
			return
		}	

		playCardPattern ,isValidated = player.playerCardsHandler.PlayCardValidateWithTopCard(playCards, topCardPattern, round)
	}

	// remove playCards from player, and add to Big2 topPlay
	// 要修正輸入多個 index 會錯誤
	var newHandCard []*Card
	for _, playCardIdx := range playCardsIdx{
		newHandCard = RemoveCardByIndex(player.handCards, playCardIdx)
	}
	player.handCards = newHandCard
	player.playCards = playCardPattern
	fmt.Printf("玩家 %s 打出了 %s %v\n\n",player.name, playCardPattern.GetKind(), player.playCards.GetCards())
	Turn++
}

func (player *HumanPlayer) play() ([]*Card,[]int){
	playCardsIdxString := strings.Fields(Scanner(fmt.Sprintf("你要出的牌是 (請依序 0 ~ %d ，以空白隔開可出多張) (如要 pass 則輸入 -1) :", len(player.handCards)-1)))
	
	if playCardsIdxString[0] == "-1"{
		return nil, nil
	}

	var playCards []*Card
	var playCardsIdx []int
	for _, playCardIdxString := range playCardsIdxString{
		playCardIdx, _ := strconv.Atoi(playCardIdxString)
		playCardsIdx = append(playCardsIdx, playCardIdx)
		card := player.handCards[playCardIdx]
		playCards = append(playCards, card)
	}
	return playCards, playCardsIdx
}


func (player *HumanPlayer) GetPlayCards() CardPattern{
	return player.playCards
}