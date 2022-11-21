package domain

import "fmt"

type PlayCardsHandler struct {
	IPlayCardsHandlerBase
}

func NewPlayCardsHandler(iPlayCardsHandlerBase IPlayCardsHandlerBase) *PlayCardsHandler {
	return &PlayCardsHandler{iPlayCardsHandlerBase}
}

func (handler *PlayCardsHandler) PlayCardValidateWithTopCard(cards []*Card, topCardsCardPattern ICardPattern, round int) (ICardPattern, bool) {
	if round == 1 && Turn == 1 {
		if !handler.validatePlayClubThreeCardPattern(cards) {
			fmt.Println("首回合第一輪必須打出包含梅花三的牌型")
			return nil, false
		}
	}

	if cardPattern, ok := handler.ValidatedCardPattern(cards); ok {
		if topCardsCardPattern == nil {
			return cardPattern, true
		}
		isBigger, error := cardPattern.IsBigger(topCardsCardPattern)
		if error != nil {
			fmt.Println(error)
			return nil, false
		}
		return cardPattern, isBigger
	}

	handler.IPlayCardsHandlerBase = handler.GetNextIPlayCardHandler()
	if handler.IPlayCardsHandlerBase != nil {
		return handler.PlayCardValidateWithTopCard(cards, topCardsCardPattern, round)
	}

	// fmt.Println("此牌型不合法，請再嘗試一次。")
	return nil, false
}

func (handler *PlayCardsHandler) ValidatedCardPattern(cards []*Card) (ICardPattern, bool) {
	cardPattern := handler.NewCardPattern(cards)
	_, isValidated := cardPattern.Validate(cardPattern)
	if isValidated {
		return cardPattern, true
	}
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
	GetNextIPlayCardHandler() IPlayCardsHandlerBase
	NewCardPattern([]*Card) ICardPattern
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
