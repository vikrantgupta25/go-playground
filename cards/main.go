package main

func main() {

	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"

	// this is only used to define the variable not while reassignment
	// this syntax is only valid inside the functions! not at the global scope
	// card := newCard()

	// array is of fixed length

	// slice is an array that can grow and shrink

	// slice and array must always be defined with data type and it should be identical

	// cards := deck{"Ace of Spades!", "Five of Diamonds!"}

	// cards := newDeck()

	// append doesn't modify the existing slice but returns a new one with added data
	// cards = append(cards, "New Element!")

	// iterating over the slice
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// hand, remainingCards := deal(cards, 5)

	// hand.print()

	// remainingCards.print()

	// cards.print()

	// cards.toString()

	// // cards.saveToFile("cardsLocal")

	// newCards, err := newDeckFromFile("cardsLocal")

	// fmt.Println(err)

	// newCards.print()

	cards := newDeck()

	cards.shuffle()

	cards.print()
}

// slice[startIndexIncluding:endingIndexNotIncluding]

// can leave numbers either side of column and goes till the end of that side

// func newCard() string {
// 	newCard := "Ace of Spades"

// 	return newCard
// }
