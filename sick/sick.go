// https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// probability returns true 1/n of times
func probability(n int) bool {
	return rand.Intn(n) == 1
}

func simulate(n int) float64 {
	numSick, numDiagnosed := 0, 0
	for i := 0; i < n; i++ {
		sick := probability(1000)
		if sick {
			numSick++
			numDiagnosed++
		} else {
			falsePositive := probability(20)
			if falsePositive {
				numDiagnosed++
			}
		}
	}

	return float64(numSick) / float64(numDiagnosed)
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(simulate(1_000_000))
	}
}
