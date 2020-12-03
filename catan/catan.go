package main

import (
	"fmt"
	"math/rand"
	"time"
)

// roll simulate a cube roll
func roll() int {
	// Intn returns values in the range [0-6), we want [1-6]
	return rand.Intn(6) + 1
}

// simulate run n simulation of two game cube rolls and returns the precentage for each total of first and second roll
func simulate(n int) map[int]float64 {
	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		val := roll() + roll()
		counts[val]++
	}

	// Convert counts to fractions
	fracs := make(map[int]float64)
	for val, count := range counts {
		frac := float64(count) / float64(n)
		fracs[val] = frac
	}
	return fracs
}

func main() {
	rand.Seed(time.Now().Unix())

	fracs := simulate(1_000_000)
	for i := 2; i <= 12; i++ {
		fmt.Printf("%2d -> %.2f\n", i, fracs[i])
	}
}
