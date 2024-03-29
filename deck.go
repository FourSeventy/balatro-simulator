package main

import (
	"math/rand"
	"time"
)

// Card represents a playing card with a value (2-14, where 11-14 are J, Q, K, A respectively) and a suit.
type Card struct {
	Value int
	Suit  string
}

// NewBaseDeck creates and returns a customized deck of cards aimed at optimizing for straights.
func NewBaseDeck() []Card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14} // 11-14 represent J, Q, K, A
	var deck []Card

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{
				Value: value,
				Suit:  suit,
			})
		}
	}

	return deck
}

// Shuffle shuffles the given deck of cards.
func Shuffle(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

// DrawHand draws a hand of n cards from the deck.
func DrawHand(deck []Card, n int) []Card {
	if n > len(deck) {
		n = len(deck)
	}
	return deck[:n]
}
