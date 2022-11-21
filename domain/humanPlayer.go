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
	playCards ICardPattern
}

func NewHumanPlayer(name string, id int, playerCardsHandler PlayCardsHandler) *HumanPlayer {
	var emptyHandCards []*Card
	humanPlayer := &HumanPlayer{name :name, id:id, handCards:emptyHandCards, playerCardsHandler:playerCardsHandler}
	return humanPlayer
}

func (player *HumanPlayer) TakeTurn(topCardPattern ICardPattern, round int) {
	fmt.Printf("輪到%s了\n",player.name)
	fmt.Println(player.name,"的手牌 :", player.handCards)

	var playCardsIdx []int
	isValidated := false
	var playCardPattern ICardPattern
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
		player.SetPlayCardsHandler(handler)
		if !isValidated{
			fmt.Println("此牌型不合法，請再嘗試一次。")
		}
	}

	// remove playCards from player, and add to Big2 topPlay
	player.removePlayCardsFromPlayer(playCardsIdx, playCardPattern)
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

func (player *HumanPlayer) removePlayCardsFromPlayer(playCardsIdx []int, playCardPattern ICardPattern) {
	newHandCards := RemoveCardsByIdx(player.handCards, playCardsIdx...)
	player.handCards = newHandCards
	player.playCards = playCardPattern
	fmt.Printf("玩家 %s 打出了 %s %v\n\n", player.name, playCardPattern.GetKind(), player.playCards.GetCards())
	Turn++
}

func (player *HumanPlayer) GetPlayCards() ICardPattern{
	return player.playCards
}

func (player *HumanPlayer) SetPlayCardsHandler(handler PlayCardsHandler){
	player.playerCardsHandler = handler
}

func (player *HumanPlayer) GetPlayerHandCards() []*Card{
	return player.handCards
}