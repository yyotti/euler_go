package main

import "testing"

var p010ATests = []struct {
	input    uint
	expected uint
}{
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: 2},
	{input: 3, expected: 2 + 3},
	{input: 4, expected: 2 + 3},
	{input: 5, expected: 2 + 3 + 5},
	{input: 10, expected: 2 + 3 + 5 + 7},
}

func TestP010A(t *testing.T) {
	for _, tt := range p010ATests {
		actual := p010A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP010B(t *testing.T) {
	for _, tt := range p010ATests {
		actual := p010B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP010A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p010A(10000)
	}
}

func BenchmarkP010B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p010B(10000)
	}
}

func ExampleP010() {
	main()

	// Output:
	// P010A: 142913828922
	// P010B: 142913828922
}
