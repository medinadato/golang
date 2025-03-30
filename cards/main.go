package main

var mynumber int

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of spades"
	// card = "Five of Diamonds"

	// card := newCard()

	// mynumber = 5
	// fmt.Println(card)
	// fmt.Println(mynumber)

	// cards := []string{newCard(), "Ace of Diamonds"}
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
