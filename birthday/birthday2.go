// Package main simulate the "Birday problem".
// See https://en.wikipedia.org/wiki/Birthday_problem for a description of the problem.
package main

import (
	"fmt"
	"math/rand"
)

// randomBirthdays returns birthdays of random group of people.
func randomBirthdays(groupSize int) []int {
	const daysInYear = 365

	birthdays := make([]int, groupSize)
	for i := 0; i < groupSize; i++ {
		birthdays[i] = rand.Intn(daysInYear)
	}

	return birthdays
}

// hasDuplicates returns true if there's at least one dumplicate number in values.
func hasDuplicates(values []int) bool {
	seen := make(map[int]bool)
	for _, n := range values {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}

// simulateBirthdays returns the fraction of groups that have at two people with the same birthday.
func simulateBirthdays(groupSize, n int) float64 {
	same := 0
	for i := 0; i < n; i++ {
		birthdays := randomBirthdays(groupSize)
		if hasDuplicates(birthdays) {
			same++
		}
	}

	return float64(same) / float64(n)
}

func main() {
	fmt.Println(simulateBirthdays(23, 1_000_00))
}
