package main

import "testing"

var p032Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: 0},
	{input: 3, expected: 0},
	{input: 4, expected: 12},
	{input: 5, expected: 52},
}

func TestP032A(t *testing.T) {
	for _, tt := range p032Tests {
		actual := p032A(tt.input)
		if actual != tt.expected {
			t.Errorf("%v: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP032B(t *testing.T) {
	for _, tt := range p032Tests {
		actual := p032B(tt.input)
		if actual != tt.expected {
			t.Errorf("%v: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP032A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p032A(7)
	}
}

func BenchmarkP032B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p032B(7)
	}
}

func ExampleP032() {
	main()

	// Output:
	// P032A: 45228
	// P032B: 45228
}
