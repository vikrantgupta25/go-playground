package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of deck
// which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Spades", "Hearts", "Diamomds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// function with receiver
// deck here is the type reference we want attach to
// the variable d is the actual copy of the variable where it is called
// the receiver is generally a single letter word in GO (standard go)
// function receivers are very similar to js prototype override
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// multiple return values from same function
func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingDeck := d[handSize:]
	return hand, remainingDeck
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// byte slice is a computer friendly way to represent a string
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		// #Option 1 log the error and return the call to newDeck()
		// #Option 2 log the error and entirely quit the program
		fmt.Println("Error:", err)
		// platform independent os functionalities
		os.Exit(1)
	}

	cards := deck(strings.Split(string(byteSlice), ","))

	return cards, nil
}

// shuffle the cards

func (d deck) shuffle() {
	for i := range d {
		randomNumber := rand.Intn(len(d) - 1)
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}
