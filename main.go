package main

import "fmt"

// mask that takes out the least significant 16 bits
const filter = 0xFFFF

// a struct that defines a generator
type generator struct {
	factor, value uint64 // both factor and value are initialized as unsigned integers,
}

// function that returns the next value of an generator
func (g *generator) NextValue() uint64 {
	g.value = (g.value * g.factor) % 2147483647
	return g.value
}

// function that count the number of matching pairs in a set of n pairs
func countMatches(n int, a, b generator) int {
	var count int // starts count as 0

	// loop through the values generated
	for i := 0; i < n; i++ {
		// if the least significant bits match, increase count
		if a.NextValue()&filter == b.NextValue()&filter {
			count++
		}
	}

	return count
}

func main() {
	a := generator{16807, 65}
	b := generator{48271, 8921}
	result := countMatches(40000000, a, b)

	fmt.Println("Number of matching pairs: ", result)
}
