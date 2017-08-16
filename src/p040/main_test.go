package main

import "testing"

var p040Tests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: -1},
	{input: 0, expected: 1},
	{input: 1, expected: 1},
	{input: 2, expected: 1},
	{input: 3, expected: 5},
	{input: 4, expected: 15},
}

func TestP040A(t *testing.T) {
	for _, tt := range p040Tests {
		actual := p040A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP040B(t *testing.T) {
	for _, tt := range p040Tests {
		actual := p040B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP040A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p040A(7)
	}
}

func BenchmarkP040B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p040B(7)
	}
}

var dTests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: -1},
	{input: 0, expected: 0},
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 10, expected: 1},
	{input: 11, expected: 0},
	{input: 12, expected: 1},
	{input: 200, expected: 0},
	{input: 210, expected: 6},
}

func TestD(t *testing.T) {
	for _, tt := range dTests {
		actual := d(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func ExampleP040() {
	main()

	// Output:
	// P040A: 210
	// P040B: 210
}
