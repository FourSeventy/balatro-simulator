package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Hill_climbing
// func HillClimb() {
// Pseudocode
// algorithm Discrete Space Hill Climbing is
//     currentNode := startNode
//     loop do
//         L := NEIGHBORS(currentNode)
//         nextEval := −INF
//         nextNode := NULL
//         for all x in L do
//             if EVAL(x) > nextEval then
//                 nextNode := x
//                 nextEval := EVAL(x)
//         if nextEval ≤ EVAL(currentNode) then
//             // Return current node since no better neighbors exist
//             return currentNode
//         currentNode := nextNode

// }

func main() {
	// deck := NewDeck()
	// Shuffle(deck)
	// hand := DrawHand(deck, 8)          // Draw a hand of 8 cards
	// straight := ContainsStraight(hand) //see if hand contains a straight
	// // print out hand
	// for _, card := range hand {
	// 	fmt.Printf("%d of %s, ", card.Value, card.Suit)
	// }

	// creates a new deck
	// deck := NewBaseDeck()

	//calculates the odds that in any random hand of 8 cards it contains a pair
	// result := MonteCarloSimulation(deck, 8, ContainsPair, 20000)
	// fmt.Printf("This deck has a %v probability to draw a pair in any given hand of size %v \n", result, 8)

	////calculates the odds that in any random hand of 8 cards it contains two pair
	//result = MonteCarloSimulation(deck, 8, ContainsTwoPair, 20000)
	//fmt.Printf("This deck has a %v probability to draw a two pair in any given hand of size %v \n", result, 8)

	////calculates the odds that in any random hand of 8 cards it contains three of a kind
	//result = MonteCarloSimulation(deck, 8, ContainsThreeOfAKind, 20000)
	//fmt.Printf("This deck has a %v probability to draw a three of a kind in any given hand of size %v \n", result, 8)

	////calculates the odds that in any random hand of 8 cards it contains a straight
	//result = MonteCarloSimulation(deck, 8, ContainsStraight, 20000)
	//fmt.Printf("This deck has a %v probability to draw a straight in any given hand of size %v \n", result, 8)

	////calculates the odds that in any random hand of 8 cards it contains a flush
	//result = MonteCarloSimulation(deck, 8, ContainsFlush, 20000)
	//fmt.Printf("This deck has a %v probability to draw a flush in any given hand of size %v \n", result, 8)

	////calculates the odds that in any random hand of 8 cards it contains a full house
	//result = MonteCarloSimulation(deck, 8, ContainsFullHouse, 20000)
	//fmt.Printf("This deck has a %v probability to draw a full house in any given hand of size %v \n", result, 8)

	////calculates the odds that in any random hand of 8 cards it contains four of a kind
	//result = MonteCarloSimulation(deck, 8, ContainsFourOfAKind, 20000)
	//fmt.Printf("This deck has a %v probability to draw a four of a kind in any given hand of size %v \n", result, 8)

	////calculates the odds that in any random hand of 8 cards it contains a straight flush
	//result = MonteCarloSimulation(deck, 8, ContainsStraightFlush, 20000)
	//fmt.Printf("This deck has a %v probability to draw a straight flush in any given hand of size %v \n", result, 8)

	deck := NewBaseDeck()
	//calculates the odds that in any random hand of 8 cards it contains a straight
	result := MonteCarloSimulation(deck, 8, ContainsStraight, 30000)
	fmt.Printf("Base Deck has a %v probability to draw a straight in any given hand of size %v \n", result, 8)

	deck = NewAbandonedDeck()
	//calculates the odds that in any random hand of 8 cards it contains a straight
	result = MonteCarloSimulation(deck, 8, ContainsStraight, 30000)
	fmt.Printf("Abandoned Deck has a %v probability to draw a straight in any given hand of size %v \n", result, 8)

}
