package main

import (
	"fmt"
)

// RunSimulation takes a deck and a condition function and tests n number of hands (shuffling each time) to determine the probability that a hand satisfies the condition
func RunSimulation(deck []Card, handSize int, conditionFunction HandConditionFunction, iterations int) float64 {
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

func main() {
	// deck := NewDeck()
	// Shuffle(deck)
	// hand := DrawHand(deck, 8)          // Draw a hand of 8 cards
	// straight := ContainsStraight(hand) //see if hand contains a straight
	// // print out hand
	// for _, card := range hand {
	// 	fmt.Printf("%d of %s, ", card.Value, card.Suit)
	// }

	//creates a new deck
	deck := NewDeck()
	//calculates the odds that in any random hand of 8 cards it contains a straight
	result := RunSimulation(deck, 8, ContainsStraight, 50000)
	fmt.Printf("This deck has a %v probability to draw a straight in any given hand of size %v \n", result, 8)

	//calculates the odds that in any random hand of 8 cards it contains a pair
	result = RunSimulation(deck, 8, ContainsPair, 50000)
	fmt.Printf("This deck has a %v probability to draw a pair in any given hand of size %v \n", result, 8)

	//calculates the odds that in any random hand of 8 cards it contains two pair
	result = RunSimulation(deck, 8, ContainsTwoPair, 50000)
	fmt.Printf("This deck has a %v probability to draw a two pair in any given hand of size %v \n", result, 8)

	//calculates the odds that in any random hand of 8 cards it contains a full house
	result = RunSimulation(deck, 8, ContainsFullHouse, 50000)
	fmt.Printf("This deck has a %v probability to draw a two pair in any given hand of size %v \n", result, 8)

}
