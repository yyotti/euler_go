package main

import "testing"

var dTests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: 1},
	{input: 3, expected: 1},
	{input: 4, expected: 3},
	{input: 5, expected: 1},
	{input: 6, expected: 6},
	{input: 7, expected: 1},
	{input: 8, expected: 7},
	{input: 9, expected: 4},
	{input: 10, expected: 8},
	{input: 220, expected: 284},
	{input: 284, expected: 220},
}

func TestD(t *testing.T) {
	for _, tt := range dTests {
		actual := d(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var p021Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: 0},
	{input: 500, expected: 504}, // (220,284)
	{input: 1000, expected: 504},
	{input: 1500, expected: 2898}, // (220,284),(1184,1210)
}

func TestP021A(t *testing.T) {
	for _, tt := range p021Tests {
		actual := p021A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP021B(t *testing.T) {
	for _, tt := range p021Tests {
		actual := p021B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP021A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p021A(5000)
	}
}

func BenchmarkP021B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p021B(5000)
	}
}

func ExampleP021() {
	main()

	// Output:
	// P021A: 31626
	// P021B: 31626
}
