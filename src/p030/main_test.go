package main

import "testing"

var powTests = []struct {
	inN      int
	inK      int
	expected int
}{
	{inN: 0, inK: 0, expected: 1},
	{inN: 0, inK: 1, expected: 0},
	{inN: 0, inK: 2, expected: 0},
	{inN: 1, inK: 0, expected: 1},
	{inN: 1, inK: 1, expected: 1},
	{inN: 1, inK: 2, expected: 1},
	{inN: 2, inK: 0, expected: 1},
	{inN: 2, inK: 1, expected: 2},
	{inN: 2, inK: 2, expected: 4},
	{inN: 2, inK: 3, expected: 8},
	{inN: 2, inK: 4, expected: 16},
	{inN: 3, inK: 0, expected: 1},
	{inN: 3, inK: 1, expected: 3},
	{inN: 3, inK: 2, expected: 9},
	{inN: 3, inK: 3, expected: 27},
	{inN: 3, inK: 4, expected: 81},
}

func TestPow(t *testing.T) {
	for _, tt := range powTests {
		actual := pow(tt.inN, tt.inK)
		if actual != tt.expected {
			t.Errorf("%d^%d: Expected %d but got %d", tt.inN, tt.inK, tt.expected, actual)
		}
	}
}

var p030Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 44},
	{input: 2, expected: 0},
	{input: 4, expected: 19316},
}

func TestP030A(t *testing.T) {
	for _, tt := range p030Tests {
		actual := p030A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP030B(t *testing.T) {
	for _, tt := range p030Tests {
		actual := p030B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP030A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p030A(4)
	}
}

func BenchmarkP030B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p030B(4)
	}
}

func ExampleP030() {
	main()

	// Output:
	// P030A: 443839
	// P030B: 443839
}
