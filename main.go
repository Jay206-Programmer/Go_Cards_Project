package main

// import "fmt"

func main() {
	// cards.saveToFile("myDeck.txt")
	// card := newDeckFromFile("myDeck.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
