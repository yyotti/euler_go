package main

import "testing"

var p028Tests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: -1},
	{input: 0, expected: -1},
	{input: 1, expected: 1},
	{input: 2, expected: -1},
	{input: 3, expected: 25},
	{input: 4, expected: -1},
	{input: 5, expected: 101},
	{input: 6, expected: -1},
	{input: 7, expected: 261},
}

func TestP028A(t *testing.T) {
	for _, tt := range p028Tests {
		actual := p028A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP028B(t *testing.T) {
	for _, tt := range p028Tests {
		actual := p028B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP028C(t *testing.T) {
	for _, tt := range p028Tests {
		actual := p028C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP018A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p028A(101)
	}
}

func BenchmarkP018B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p028B(101)
	}
}

func BenchmarkP018C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p028C(101)
	}
}

var tTests = []struct {
	input    int
	expected int
}{
	{input: 1, expected: 1},
	{input: 3, expected: 24},
	{input: 5, expected: 76},
	{input: 7, expected: 160},
}

func TestT(te *testing.T) {
	for _, tt := range tTests {
		actual := t(tt.input)
		if actual != tt.expected {
			te.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func ExampleP028() {
	main()

	// Output:
	// P028A: 669171001
	// P028B: 669171001
	// P028C: 669171001
}
