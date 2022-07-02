package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	d = d[handSize:]

	return hand, d
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		os.Exit(1)
	}

	deckString := string(data)

	return deck(strings.Split(deckString, ","))
}

func (d deck) shuffle() deck {
	t := time.Now().UnixNano()
	s := rand.NewSource(t)
	r := rand.New(s)

	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}

	return d
}