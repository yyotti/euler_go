package main

import "testing"

var isAbundantTests = []struct {
	input    int
	expected bool
}{
	{input: -1, expected: false},
	{input: 0, expected: false},
	{input: 1, expected: false},
	{input: 2, expected: false},
	{input: 3, expected: false},
	{input: 4, expected: false},
	{input: 5, expected: false},
	{input: 6, expected: false},
	{input: 7, expected: false},
	{input: 8, expected: false},
	{input: 9, expected: false},
	{input: 10, expected: false},
	{input: 11, expected: false},
	{input: 12, expected: true},
	{input: 13, expected: false},
	{input: 14, expected: false},
	{input: 15, expected: false},
	{input: 16, expected: false},
	{input: 17, expected: false},
	{input: 18, expected: true},
	{input: 19, expected: false},
	{input: 20, expected: true},
}

func TestIsAbundnt(t *testing.T) {
	for _, tt := range isAbundantTests {
		actual := isAbundant(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

const p23Expected = 4179871

func TestP023A(t *testing.T) {
	actual := p023A()
	if actual != p23Expected {
		t.Errorf("Expected %v but got %v", p23Expected, actual)
	}
}

func TestP023B(t *testing.T) {
	actual := p023B()
	if actual != p23Expected {
		t.Errorf("Expected %v but got %v", p23Expected, actual)
	}
}

func BenchmarkP023A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p023A()
	}
}

func BenchmarkP023B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p023B()
	}
}

func ExampleP023() {
	main()

	// Output:
	// P023A: 4179871
	// P023B: 4179871
}
