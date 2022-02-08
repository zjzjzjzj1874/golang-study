package main

import "fmt"

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "five of hearts")

	for i, card := range cards {
		fmt.Printf("%d,%s\n", i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

type deck []string
