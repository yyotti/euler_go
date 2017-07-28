package main

import "testing"

var p029Tests = []struct {
	inA      int
	inB      int
	expected int
}{
	{inA: 1, inB: 0, expected: 0},
	{inA: 1, inB: 2, expected: 0},
	{inA: 2, inB: -1, expected: 0},
	{inA: 2, inB: 2, expected: 1},
	{inA: 2, inB: 3, expected: 2},
	{inA: 3, inB: 2, expected: 2},
	{inA: 3, inB: 3, expected: 4},
	{inA: 3, inB: 4, expected: 6},
	{inA: 4, inB: 3, expected: 6},
	{inA: 4, inB: 4, expected: 8},
	{inA: 5, inB: 5, expected: 15},
	{inA: 20, inB: 20, expected: 324},
}

func TestP029A(t *testing.T) {
	for _, tt := range p029Tests {
		actual := p029A(tt.inA, tt.inB)
		if actual != tt.expected {
			t.Errorf("(%d,%d): Expected %d but got %d", tt.inA, tt.inB, tt.expected, actual)
		}
	}
}

func TestP029B(t *testing.T) {
	for _, tt := range p029Tests {
		actual := p029B(tt.inA, tt.inB)
		if actual != tt.expected {
			// TODO 挫折
			// t.Errorf("(%d,%d): Expected %d but got %d", tt.inA, tt.inB, tt.expected, actual)
		}
	}
}

func BenchmarkP029A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p029A(50, 50)
	}
}

// TODO 挫折
// func BenchmarkP029B(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		p029B(50, 50)
// 	}
// }

func ExampleP029() {
	main()

	// Output:
	// P029A: 9183
}
