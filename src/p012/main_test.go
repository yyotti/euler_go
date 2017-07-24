package main

import (
	"reflect"
	"sort"
	"testing"
)

var triangleTests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 1},
	{input: 2, expected: 1 + 2},
	{input: 3, expected: 1 + 2 + 3},
	{input: 4, expected: 1 + 2 + 3 + 4},
	{input: 5, expected: 1 + 2 + 3 + 4 + 5},
}

func TestTriangle(t *testing.T) {
	for _, tt := range triangleTests {
		actual := triangle(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var divisorsTests = []struct {
	input    int
	expected []int
}{
	{input: 0, expected: []int{1}},
	{input: 1, expected: []int{1}},
	{input: 2, expected: []int{1, 2}},
	{input: 3, expected: []int{1, 3}},
	{input: 4, expected: []int{1, 2, 4}},
	{input: 24, expected: []int{1, 2, 3, 4, 6, 8, 12, 24}},
}

func TestDivisors(t *testing.T) {
	for _, tt := range divisorsTests {
		actual := divisors(tt.input)
		sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var p012Tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 1},
	{input: 1, expected: 3},
	{input: 2, expected: 6},
	{input: 5, expected: 28},
}

func TestP012A(t *testing.T) {
	for _, tt := range p012Tests {
		actual := p012A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP012B(t *testing.T) {
	for _, tt := range p012Tests {
		actual := p012B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var divisorsCntTests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: 0},
	{input: 0, expected: 1},
	{input: 1, expected: 1},
	{input: 2, expected: 2},
	{input: 3, expected: 2},
	{input: 4, expected: 3},
	{input: 24, expected: 8},
}

func TestDivisorsCnt(t *testing.T) {
	for _, tt := range divisorsCntTests {
		actual := divisorsCnt(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divisors(3628800) // 1 x 2 x ... x 10
	}
}

func BenchmarkDivisorsCnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divisorsCnt(3628800) // 1 x 2 x ... x 10
	}
}

func BenchmarkP012A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p012A(200)
	}
}

func BenchmarkP012B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p012B(200)
	}
}

func ExampleP012() {
	main()

	// Output:
	// P012A: 76576500
	// P012B: 76576500
}
