package main

import "testing"

var p005Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 3, expected: 6},
	{input: 4, expected: 12},
	{input: 5, expected: 60},
	{input: 6, expected: 60},
	{input: 7, expected: 420},
	{input: 8, expected: 840},
	{input: 9, expected: 2520},
	{input: 10, expected: 2520},
}

func TestP005A(t *testing.T) {
	for _, tt := range p005Tests {
		actual := p005A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func TestP005B(t *testing.T) {
	for _, tt := range p005Tests {
		actual := p005B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP005A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p005A(10)
	}
}

func BenchmarkP005B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p005B(10)
	}
}

func ExampleP005() {
	main()

	// Output:
	// P005A: 232792560
	// P005B: 232792560
}
