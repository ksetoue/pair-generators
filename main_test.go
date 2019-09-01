package main

import "testing"

func TestGenerators(t *testing.T) {
	tests := []struct {
		g      generator
		expect []uint64
	}{
		{generator{16807, 65}, []uint64{1092455, 1181022009, 245556042, 1744312007, 1352636452}},    // first values generated for A
		{generator{48271, 8921}, []uint64{430625591, 1233683848, 1431495498, 137874439, 285222916}}, //first values generated for B
	}

	for _, test := range tests {
		leng := len(test.expect)
		for i := 0; i < leng; i++ {
			result := test.g.NextValue()
			if result != test.expect[i] {
				t.Errorf("Expected %d, returned value was %d", test.expect[i], result)
			}
		}
	}
}

func TestCountMatches(t *testing.T) {
	tests := []struct {
		initialValueA, initialValueB uint64
		Expected                     int
	}{
		{65, 8921, 588},
	}

	for _, test := range tests {
		a := generator{16807, test.initialValueA}
		b := generator{48271, test.initialValueB}

		result := CountMatches(40000000, a, b)
		if result != test.Expected {
			t.Errorf("Expected %d, returned value was %d", test.Expected, result)
		}
	}
}
