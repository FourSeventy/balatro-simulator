package main

import (
	"gonum.org/v1/gonum/stat/combin"
)

// calculatePairOdds takes a deck and a handsize and calculates the odds of drawing a pair using combinatorics
// TODO: WIP
func calculatePairOdds(deck []Card, handSize int) {
	//what we want to compute is the number of possible pairs, and divide it by the total number of hands.

	//total number of hands is: deckSize choose handSize or deckSize! / handSize!(deckSize - handSize)!
	combin.Binomial(len(deck), handSize)

}
