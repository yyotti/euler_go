package main

import (
	"reflect"
	"testing"
)

var recTests = []struct {
	input    int
	expected []int
}{
	{input: -1, expected: nil},
	{input: 0, expected: nil},
	{input: 1, expected: []int{}},
	{input: 2, expected: []int{}},
	{input: 3, expected: []int{3}},
	{input: 4, expected: []int{}},
	{input: 5, expected: []int{}},
	{input: 6, expected: []int{6}},
	{input: 7, expected: []int{1, 4, 2, 8, 5, 7}},
	{input: 8, expected: []int{}},
	{input: 9, expected: []int{1}},
	{input: 10, expected: []int{}},
	{input: 11, expected: []int{0, 9}},
	{input: 12, expected: []int{3}},
	{input: 13, expected: []int{0, 7, 6, 9, 2, 3}},
}

func TestRec(t *testing.T) {
	for _, tt := range recTests {
		actual := rec(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var p026Tests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: 0},
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 2, expected: 0},
	{input: 10, expected: 7},
	{input: 20, expected: 19},
	{input: 30, expected: 29},
	{input: 40, expected: 29},
	{input: 50, expected: 47},
	{input: 60, expected: 59},
	{input: 70, expected: 61},
	{input: 80, expected: 61},
	{input: 90, expected: 61},
	{input: 100, expected: 97},
}

func TestP026A(t *testing.T) {
	for _, tt := range p026Tests {
		actual := p026A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP026B(t *testing.T) {
	for _, tt := range p026Tests {
		actual := p026B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP026C(t *testing.T) {
	for _, tt := range p026Tests {
		actual := p026C(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP026D(t *testing.T) {
	for _, tt := range p026Tests {
		actual := p026D(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP026A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p026A(100)
	}
}

func BenchmarkP026B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p026B(100)
	}
}

func BenchmarkP026C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p026C(100)
	}
}

func BenchmarkP026D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p026D(100)
	}
}

func ExampleP026() {
	main()

	// Output:
	// P026A: 983
	// P026B: 983
	// P026C: 983
	// P026D: 983
}
