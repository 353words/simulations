// Package main simulate percision of a medical test.
// See https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/
// for more details.
package main

import (
	"fmt"
	"math/rand"
)

// probability returns true 1/n of times.
func probability(n int) bool {
	return rand.Intn(n) == 1
}

// simulate run selects n random people and return the fraction of people
// actually sick from the total number of people diagnosed as sick.
func simulate(n int) float64 {
	numSick, numDiagnosed := 0, 0
	for i := 0; i < n; i++ {
		sick := probability(1000) // A person has 1/1000 chance of being sick.
		if sick {
			numSick++
			numDiagnosed++
		} else {
			// A healthy person has a 5% change of being diagnosed as sick.
			falsePositive := probability(20)
			if falsePositive {
				numDiagnosed++
			}
		}
	}

	return float64(numSick) / float64(numDiagnosed)
}

func main() {
	fmt.Println(simulate(1_000_000))
}
