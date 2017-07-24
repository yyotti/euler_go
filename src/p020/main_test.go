package main

import "testing"

var p020Tests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: 0},
	{input: 0, expected: 1},
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 3, expected: 6},
	{input: 4, expected: 6},
	{input: 5, expected: 3},
	{input: 6, expected: 9},
	{input: 7, expected: 9},
	{input: 8, expected: 9},
	{input: 9, expected: 27},
	{input: 10, expected: 27},
}

func TestP020A(t *testing.T) {
	for _, tt := range p020Tests {
		actual := p020A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP020B(t *testing.T) {
	for _, tt := range p020Tests {
		actual := p020B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP020A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p020A(30)
	}
}

func BenchmarkP020B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p020B(30)
	}
}

func ExampleP020() {
	main()

	// Output:
	// P020A: 648
	// P020B: 648
}
