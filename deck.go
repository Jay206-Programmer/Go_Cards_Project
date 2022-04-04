package main

import "fmt"

// * Create a new type of 'deck'
// * Which is a slice of strings

type deck []string

/*
	Get a new Deck.
*/
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four",
							"Five", "Six", "Seven", "Eight",
							"Nine", "Ten", "Jack", "King", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+ " of "+suit)
		}
	}

	return cards
}

/* 
	Used to Print All the cards inside the Deck.
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
