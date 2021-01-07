// Package main simulate percision of a medical test.
// See https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/
// for more details.
package main

import (
	"fmt"
	"math/rand"
)

// oneChanceIn returns true one in n times.
func oneChanceIn(n int) bool {
	return rand.Intn(n) == 1
}

// simulate run selects sampleSize random people and return the fraction of people
// actually sick from the total number of people diagnosed as sick.
func simulate(sampleSize int) float64 {
	var numSick, numDiagnosed int
	for i := 0; i < sampleSize; i++ {
		// A person has 1/1000 chance of being sick.
		sick := oneChanceIn(1000)
		if sick {
			numSick++
			numDiagnosed++
		} else {
			// A healthy person has a 5% chance of being diagnosed as sick.
			if oneChanceIn(20) {
				numDiagnosed++
			}
		}
	}

	return float64(numSick) / float64(numDiagnosed)
}

func main() {
	fmt.Println(simulate(1_000_000))
}
