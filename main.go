package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of spades")
	// cards.print()
	// first, second := cards.deal(2)

	// fmt.Println(first)
	// fmt.Println(second)
	// fmt.Println(cards.toString())
	// _ = cards.saveToFile("cards.txt")

	cards := newDeckFromFile("cards.txt")

	cards.print()
	cards.shuffle()
	cards.print()
}
