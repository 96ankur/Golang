package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string // Create a new type of 'deck' which is a slice of strings

/**
	RECEIVER FUNCTION (similar to class)
	Any variable of type deck now gets access to the "print" method.
	The actual copy of the deck we are working with is available in the
	function as a variable called 'd'
*/

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // syntax to create a slice (Dynamic length array)
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // _ means that there is a variable but we don't want to use it
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // add an element to the last of the silce and return
			// the same array
		}
	}

	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]   // swapping of two variable
	}
}