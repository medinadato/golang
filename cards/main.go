package main

// var mynumber int
import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of spades"
	// card = "Five of Diamonds"

	// card := newCard()

	// mynumber = 5
	// fmt.Println(card)
	// fmt.Println(mynumber)

	// cards := []string{newCard(), "Ace of Diamonds"}
	// cards := deck{newCard(), "Ace of Diamonds"}
	// cards = append(cards, "Six of spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards.saveToFile("cards/var/my_cards.txt")

	cards := newDeckFromFile("cards/var/my_cards.txt")
	cards.shuffle()
	fmt.Println(cards.toString())

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
