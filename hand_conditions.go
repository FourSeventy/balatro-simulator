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

// ContainsPair checks if the hand contains a poker pair.
func ContainsPair(hand []Card) bool {
	valueCount := make(map[int]int) // Map to count occurrences of each card value.
	for _, card := range hand {
		valueCount[card.Value]++
		if valueCount[card.Value] == 2 { // If any card value occurs twice, we have a pair.
			return true
		}
	}
	return false
}

// ContainsTwoPair checks if the hand contains two pairs.
func ContainsTwoPair(hand []Card) bool {
	valueCount := make(map[int]int) // Map to count occurrences of each card value.
	pairCount := 0

	for _, card := range hand {
		valueCount[card.Value]++
		// If we've found a pair, increase the pair count. Then reset the count for that value to avoid counting it as 2 pairs if we see it again.
		if valueCount[card.Value] == 2 {
			pairCount++
			valueCount[card.Value] = 0 // Reset to avoid double counting this value as two pairs.
		}
	}
	// If we've found exactly two pairs, return true.
	return pairCount >= 2
}

//=========== Helper Functions =================

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
