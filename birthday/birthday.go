// https://en.wikipedia.org/wiki/Birthday_problem
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// hasSame returns True is a random group has the same birthday
func hasSame(groupSize int) bool {
	const daysInYear = 365

	seen := make(map[int]bool)
	for i := 0; i < groupSize; i++ {
		day := rand.Intn(daysInYear)
		if seen[day] {
			return true
		}
		seen[day] = true
	}

	return false
}

// simulateBirthdays returns the fraction of groups that have at two people
// with the same birthday
func simulateBirthdays(groupSize, n int) float64 {
	same := 0
	for i := 0; i < n; i++ {
		if hasSame(groupSize) {
			same++
		}
	}

	return float64(same) / float64(n)
}

func main() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		fmt.Println(simulateBirthdays(23, 1_000_00))
	}
}
