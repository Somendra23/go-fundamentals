package main

import (
	"fmt"
)

var deckSize string = "20"

// cant do this :cannot initialize outside func in this fashion
// deckSize:=20

func newCard() string {
	return "King of Diamonds"
}

func handleArraySlice() {
	//Array and Slice are fixed and dynamic size containers
	cards := deck{"Ace of Spades", newCard()}
	//cardsUpdated := []string{}
	cardsUpdated := deck{}
	cardsUpdated = append(cards, "Jack of Heart")
	//original slice data is not altered by append function
	//fmt.Println(cards)
	// fmt.Println(cardsUpdated)
	cardsUpdated.printCards()
	cards.printCards()
	// for i, card := range cards {
	// 	fmt.Println("Iterating over cards : ")
	// 	fmt.Println(i, card)
	// }

	colors := []string{"Red", "Green", "Blue"}
	for index, color := range colors {
		fmt.Println(index, color)
	}

}

func main() {
	//var card string = "Ace of Spades"
	// handleArraySlice()
	// card := "Ace of Spades"
	// card = newCard()
	// fmt.Println(card)
	// fmt.Println(deckSize)

	//the below is not intialized so will give error
	//rateLimit = 15

	rateLimit := 15

	fmt.Println(rateLimit)
	//card := newDeck()
	//card.printCards()
	//hand, stacked := deal(card, 5)
	//fmt.Println(hand)
	//fmt.Println("=======================")
	//fmt.Println(stacked)
	//hand.saveToFile("my_cards")
	//fmt.Println("read from my_cards====")
	//cards := newDeckFromFile("my_cards")
	//cards.printCards()

	//shuffle example
	cards := newDeck()
	cards.shuffle()
	cards.printCards()
}
