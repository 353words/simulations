// Package main simulate the "Birday problem".
// See https://en.wikipedia.org/wiki/Birthday_problem for a description of the problem.
package main

import (
	"fmt"
	"math/rand"
)

// duplicateBirthdayInGroup returns True if at least two people in a random
// group has the same birthday.
func duplicateBirthdayInGroup(groupSize int) bool {
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

// simulateBirthdays draws "n" random group of "groupSize" people and returns
// the fraction of groups that have at least two people with the same birthday.
func simulateBirthdays(groupSize, n int) float64 {
	same := 0
	for i := 0; i < n; i++ {
		if duplicateBirthdayInGroup(groupSize) {
			same++
		}
	}

	return float64(same) / float64(n)
}

func main() {
	fmt.Println(simulateBirthdays(23, 1_000_00))
}
