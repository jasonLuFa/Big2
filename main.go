package main

import (
	"big2/domain"
	playCardsHandler "big2/domain/playCardsHandler"
)

func main() {

	b:= domain.NewBig2()
	b.Start(*domain.NewPlayCardsHandler(
		playCardsHandler.NewSingleHandler(
		playCardsHandler.NewPairHandler(
		playCardsHandler.NewFullHouseHandler(
		playCardsHandler.NewStraightHandler(
		nil))))))
	b.ShowWinner()
	
}