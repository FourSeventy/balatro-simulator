package main

import "sort"

// this type represents a function that tests if a given hand satisfies a condition
type HandConditionFunction func([]Card) bool

// MonteCarloSimulation takes a deck and a condition function and tests n number of hands (shuffling each time)
// to determine the probability that a hand satisfies the condition
func MonteCarloSimulation(deck []Card, handSize int, conditionFunction HandConditionFunction, iterations int) float64 {
	satisfactions := 0.0
	for i := 0; i < iterations; i++ {
		//shuffle deck
		Shuffle(deck)
		//draw hand
		hand := DrawHand(deck, handSize)
		//see if hand satisfies function
		result := conditionFunction(hand)
		if result == true {
			satisfactions += 1
		}
	}

	return satisfactions / float64(iterations)
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

// ContainsPair checks if the hand contains a three of a kind
func ContainsThreeOfAKind(hand []Card) bool {
	valueCount := make(map[int]int) // Map to count occurrences of each card value.
	for _, card := range hand {
		valueCount[card.Value]++
		if valueCount[card.Value] == 3 { // If any card value occurs three times
			return true
		}
	}
	return false
}

// ContainsPair checks if the hand contains a four of a kind
func ContainsFourOfAKind(hand []Card) bool {
	valueCount := make(map[int]int) // Map to count occurrences of each card value.
	for _, card := range hand {
		valueCount[card.Value]++
		if valueCount[card.Value] == 4 { // If any card value occurs four times
			return true
		}
	}
	return false
}

// ContainsFullHouse checks if the hand contains a full house.
func ContainsFullHouse(hand []Card) bool {
	valueCount := make(map[int]int) // Map to count occurrences of each card value.
	for _, card := range hand {
		valueCount[card.Value]++
	}

	// Variables to check for the presence of a pair and a three of a kind.
	var hasPair, hasThreeOfAKind bool
	for _, count := range valueCount {
		if count == 2 {
			hasPair = true
		} else if count == 3 {
			hasThreeOfAKind = true
		}
	}

	// A full house requires both a pair and a three of a kind.
	return hasPair && hasThreeOfAKind
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

// ContainsFlush checks if the hand contains a flush.
func ContainsFlush(hand []Card) bool {
	suitCount := make(map[string]int) // Map to count occurrences of each suit.
	for _, card := range hand {
		suitCount[card.Suit]++
		if suitCount[card.Suit] >= 5 { // If any suit occurs five or more times, we have a flush.
			return true
		}
	}
	return false
}

// ContainsStraightFlush checks if the hand contains a straight flush.
// TODO: I think this has a bug?
func ContainsStraightFlush(hand []Card) bool {
	if len(hand) < 5 {
		return false // Can't form a straight flush with less than 5 cards
	}

	sortByValueThenSuit(hand)

	for i := 0; i < len(hand)-4; i++ { // Need at least 5 cards for a straight flush
		if hand[i].Suit == hand[i+4].Suit &&
			hand[i].Value+1 == hand[i+1].Value &&
			hand[i].Value+2 == hand[i+2].Value &&
			hand[i].Value+3 == hand[i+3].Value &&
			hand[i].Value+4 == hand[i+4].Value {
			return true
		}
	}
	return false
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

// sortByValueThenSuit sorts the cards first by value, then by suit.
func sortByValueThenSuit(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Value == cards[j].Value {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Value < cards[j].Value
	})
}
