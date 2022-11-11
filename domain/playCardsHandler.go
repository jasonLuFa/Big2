package domain

import "fmt"

type PlayCardsHandler struct {
	IPlayCardsHandlerBase
}

func NewPlayCardsHandler(iPlayCardsHandlerBase IPlayCardsHandlerBase) *PlayCardsHandler {
	return &PlayCardsHandler{iPlayCardsHandlerBase}
}

func (handler *PlayCardsHandler) PlayCardValidateWithTopCard(cards []*Card, topCardsCardPattern CardPattern, round int) (CardPattern, bool) {
	if round == 1 && Turn == 1{
		if !handler.validatePlayClubThreeCardPattern(cards){
			fmt.Println("首回合第一輪必須打出包含梅花三的牌型")
			return nil, false
		}
	}

	// TODO: 這邊驗證出牌型好像出問題
	if cardPattern, ok := handler.ValidatedCardPattern(cards); ok {
		fmt.Println("topCardsCardPattern  =.+ :",topCardsCardPattern)
		if topCardsCardPattern == nil {
			fmt.Println("isValidated", true)
			return cardPattern, true
		}
		isBigger, error := cardPattern.IsBigger(topCardsCardPattern)
		if error != nil {
			fmt.Println(error)
			return nil ,false
		}
		return cardPattern, isBigger
	}

	handler.IPlayCardsHandlerBase = handler.GetNextIPlayCardHandler()
	if handler.IPlayCardsHandlerBase != nil {
		return handler.PlayCardValidateWithTopCard(cards, topCardsCardPattern, round)
	}

	fmt.Println("此牌型不合法，請再嘗試一次。")
	return nil, false
}

func (handler *PlayCardsHandler) validatePlayClubThreeCardPattern(cards []*Card) bool {
	isContainClubThree := false
	for _, card := range cards {
		if card.GetRankKind() == 0 && card.GetSuitKind() == 0 {
			isContainClubThree = true
		}
	}
	return isContainClubThree
}

type IPlayCardsHandlerBase interface {
	ValidatedCardPattern([]*Card) (CardPattern, bool)
	GetNextIPlayCardHandler() IPlayCardsHandlerBase
}

type PlayCardsHandlerBase struct {
	next IPlayCardsHandlerBase
}

func NewPlayCardsHandlerBase(next IPlayCardsHandlerBase) *PlayCardsHandlerBase {
	return &PlayCardsHandlerBase{next}
}

func (handlerBase PlayCardsHandlerBase) GetNextIPlayCardHandler() IPlayCardsHandlerBase {
	return handlerBase.next
}
