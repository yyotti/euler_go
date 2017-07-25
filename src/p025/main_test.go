package main

import "testing"

var p025Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: -1},
	{input: 1, expected: 1},
	{input: 2, expected: 7},
	{input: 3, expected: 12},
	{input: 4, expected: 17},
}

func TestP025A(t *testing.T) {
	for _, tt := range p025Tests {
		actual := p025A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP025B(t *testing.T) {
	for _, tt := range p025Tests {
		actual := p025B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP025A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p025A(100)
	}
}

func BenchmarkP025B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p025B(100)
	}
}

func ExampleP025() {
	main()

	// Output:
	// P025A: 4782
	// P025B: 4782
}
