package main

import (
	"reflect"
	"testing"
)

var p035Tests = []struct {
	input    int
	expected int
}{
	{input: -1, expected: 0},
	{input: 0, expected: 0},
	{input: 1, expected: 0},
	{input: 10, expected: 4},
	{input: 20, expected: 7},
	{input: 50, expected: 9},
	{input: 100, expected: 13},
}

func TestP035A(t *testing.T) {
	for _, tt := range p035Tests {
		actual := p035A(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func TestP035B(t *testing.T) {
	for _, tt := range p035Tests {
		actual := p035B(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkP035A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p035A(10000)
	}
}

func BenchmarkP035B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p035B(10000)
	}
}

var containsInvalidNumTests = []struct {
	input    []int
	expected bool
}{
	{input: []int{}, expected: false},
	{input: []int{0}, expected: false},
	{input: []int{1}, expected: false},
	{input: []int{1, 0}, expected: true},
	{input: []int{1, 1}, expected: false},
	{input: []int{1, 2}, expected: true},
	{input: []int{5, 1, 3}, expected: true},
	{input: []int{7, 3, 1}, expected: false},
}

func TestContainsInvalidNum(t *testing.T) {
	for _, tt := range containsInvalidNumTests {
		actual := containsInvalidNum(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var rotateTests = []struct {
	input    []int
	expected []int
}{
	{input: []int{0}, expected: []int{0}},
	{input: []int{1}, expected: []int{1}},
	{input: []int{1, 0}, expected: []int{0, 1}},
	{input: []int{1, 1}, expected: []int{1, 1}},
	{input: []int{1, 2}, expected: []int{2, 1}},
	{input: []int{1, 0, 0}, expected: []int{0, 1, 0}},
	{input: []int{1, 0, 1}, expected: []int{1, 1, 0}},
	{input: []int{1, 1, 0}, expected: []int{0, 1, 1}},
	{input: []int{2, 1, 1}, expected: []int{1, 2, 1}},
}

func TestRotate(t *testing.T) {
	for _, tt := range rotateTests {
		actual := rotate(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var rotateNTests = []struct {
	inN      int
	inR      int
	expected int
}{
	{inN: 123, inR: -1, expected: 123},
	{inN: 123, inR: 0, expected: 123},
	{inN: 1, inR: 1, expected: 1},
	{inN: 1, inR: 2, expected: 1},
	{inN: 10, inR: 1, expected: 1},
	{inN: 10, inR: 2, expected: 10},
	{inN: 10, inR: 3, expected: 1},
	{inN: 10, inR: 4, expected: 10},
	{inN: 11, inR: 1, expected: 11},
	{inN: 11, inR: 2, expected: 11},
	{inN: 11, inR: 3, expected: 11},
	{inN: 11, inR: 4, expected: 11},
	{inN: 12, inR: 1, expected: 21},
	{inN: 12, inR: 2, expected: 12},
	{inN: 12, inR: 3, expected: 21},
	{inN: 12, inR: 4, expected: 12},
	{inN: 100, inR: 1, expected: 10},
	{inN: 100, inR: 2, expected: 1},
	{inN: 100, inR: 3, expected: 100},
	{inN: 100, inR: 4, expected: 10},
	{inN: 100, inR: 5, expected: 1},
	{inN: 101, inR: 1, expected: 110},
	{inN: 101, inR: 2, expected: 11},
	{inN: 101, inR: 3, expected: 101},
	{inN: 101, inR: 4, expected: 110},
	{inN: 101, inR: 5, expected: 11},
	{inN: 110, inR: 1, expected: 11},
	{inN: 110, inR: 2, expected: 101},
	{inN: 110, inR: 3, expected: 110},
	{inN: 110, inR: 4, expected: 11},
	{inN: 110, inR: 5, expected: 101},
	{inN: 211, inR: 1, expected: 121},
	{inN: 211, inR: 2, expected: 112},
	{inN: 211, inR: 3, expected: 211},
	{inN: 211, inR: 4, expected: 121},
	{inN: 211, inR: 5, expected: 112},
}

func TestRotateN(t *testing.T) {
	for _, tt := range rotateNTests {
		actual := rotateN(tt.inN, tt.inR)
		if actual != tt.expected {
			t.Errorf("%d,%d: Expected %d but got %d", tt.inN, tt.inR, tt.expected, actual)
		}
	}
}

var joinNumsTests = []struct {
	input    []int
	expected int
}{
	{input: []int{0}, expected: 0},
	{input: []int{1}, expected: 1},
	{input: []int{1, 0}, expected: 10},
	{input: []int{1, 1}, expected: 11},
	{input: []int{1, 2}, expected: 12},
	{input: []int{1, 0, 0}, expected: 100},
	{input: []int{0, 0, 1}, expected: 1},
	{input: []int{0, 1, 0}, expected: 10},
	{input: []int{2, 1, 1}, expected: 211},
}

func TestJoinNums(t *testing.T) {
	for _, tt := range joinNumsTests {
		actual := joinNums(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

func ExampleP035() {
	main()

	// Output:
	// P035A: 55
	// P035B: 55
}
