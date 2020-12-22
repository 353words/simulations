// Package main simulate percision of a medical test.
// See https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/
// for more details.
package main

import (
	"fmt"
	"math/rand"
)

// positiveCase returns true for 1 patient in a specified group of patients.
func positiveCase(numberOfPatients int) bool {
	return rand.Intn(numberOfPatients) == 1
}

// simulate run selects n random people and return the fraction of people
// actually sick from the total number of people diagnosed as sick.
func simulate(n int) float64 {
	var numSick, numDiagnosed int
	for i := 0; i < n; i++ {
		// A person has 1/1000 chance of being sick.
		sick := positiveCase(1000)
		if sick {
			numSick++
			numDiagnosed++
			continue
		}
		// A healthy person has a 5% chance of being diagnosed as sick.
		if positiveCase(20) {
			numDiagnosed++
		}
	}

	return float64(numSick) / float64(numDiagnosed)
}

func main() {
	fmt.Println(simulate(1_000_000))
}
