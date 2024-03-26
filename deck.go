package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d *deck) print() {

	for i, card := range *d {
		fmt.Println(i, card)
	}

}

func (d *deck) deal(handSize int) (deck, deck) {
	return (*d)[:handSize], (*d)[handSize:]
}

func (d *deck) toString() string {
	return strings.Join([]string(*d), ",")
}

func (d *deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := (os.ReadFile(filename))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fileContent := string(bs)
	cardsSlice := strings.Split(fileContent, ",")

	return deck(cardsSlice)
}

func (d *deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(source)

	for i := range *d {
		newPosition := rd.Intn(len(*d) - 1)
		(*d)[i], (*d)[newPosition] = (*d)[newPosition], (*d)[i]
	}
}
