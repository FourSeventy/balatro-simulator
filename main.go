package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Card represents a playing card with a value (2-14, where 11-14 are J, Q, K, A respectively) and a suit.
type Card struct {
	Value int
	Suit  string
}

// NewDeck creates and returns a customized deck of cards aimed at optimizing for straights.
func NewDeck() []Card {
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

// ContainsStraight takes a slice of Cards and returns true if the hand contains a straight.
func ContainsStraight(hand []Card) bool {
	if len(hand) < 5 {
		return false // Cannot form a straight with fewer than 5 cards
	}

	// Create a slice of values for easier processing
	values := make([]int, len(hand))
	for i, card := range hand {
		values[i] = card.Value
	}

	// Sort the values
	sort.Ints(values)

	// Check for straights
	for i := 0; i < len(values)-4; i++ {
		if isSequential(values[i:]) {
			return true
		}
	}

	// Special case for Ace acting as the lowest card (value of 1)
	// Check if the hand contains an Ace and four cards 2-5
	if contains(values, 14) {
		// Treat Ace as 1
		values = append([]int{1}, values...)
		sort.Ints(values) // Re-sort with Ace as 1
		for i := 0; i < len(values)-4; i++ {
			if isSequential(values[i:]) {
				return true
			}
		}
	}

	return false
}

// isSequential checks if the first 5 cards in a sorted slice of card values form a straight.
func isSequential(values []int) bool {
	for i := 0; i < 4; i++ {
		if values[i]+1 != values[i+1] {
			return false
		}
	}
	return true
}

// contains checks if a slice of integers contains a specific integer.
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	deck := NewDeck()
	Shuffle(deck)
	hand := DrawHand(deck, 8)          // Draw a hand of 8 cards
	straight := ContainsStraight(hand) //see if hand contains a straight

	// For simplicity, just print out the hand for now.
	for _, card := range hand {
		fmt.Printf("%d of %s, ", card.Value, card.Suit)
	}

	fmt.Printf("\nContains straight: %v\n", straight)
}
