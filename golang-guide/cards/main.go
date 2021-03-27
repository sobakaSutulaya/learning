package main

func main() {
	// cards := newDeck()
	// cards.save("cards.txt")
	cards := newDeckFromFile("cards.txt")
    cards.print()
}
