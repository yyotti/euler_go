package main

import "testing"

const p027Expected = -59231

var p027Tests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: 0},
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: -2},
	{input: 41, expected: -41},
	{input: 42, expected: -41},
	{input: 50, expected: -235},
}

func TestP027A(t *testing.T) {
	for _, tt := range p027Tests {
		actual := p027A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP027B(t *testing.T) {
	for _, tt := range p027Tests {
		actual := p027B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP027A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p027A(max)
	}
}

func BenchmarkP027B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p027B(max)
	}
}

func ExampleP027() {
	main()

	// Output:
	// P027A: -59231
}
