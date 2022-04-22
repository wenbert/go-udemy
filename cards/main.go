package main

import "fmt"

func main() {
	//var card string = "Ace of Spades
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	//fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
