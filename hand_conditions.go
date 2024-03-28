package main

import "sort"

// this type represents a function that tests if a given hand satisfies a condition
type HandConditionFunction func([]Card) bool

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
