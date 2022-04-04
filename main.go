package main

// import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("myDeck.txt")
	card := newDeckFromFile("myDeck.txt")
	card.shuffle()
	card.print()
}
