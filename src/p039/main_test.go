package main

import "testing"

var p039Tests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: 0},
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 12, expected: 12},
	{input: 20, expected: 12},
	{input: 24, expected: 24},
	{input: 27, expected: 24},
	{input: 120, expected: 120},
}

func TestP039A(t *testing.T) {
	for _, tt := range p039Tests {
		actual := p039A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP039A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p039A(500)
	}
}

func ExampleP039() {
	main()

	// Output:
	// P039A: 840
}
