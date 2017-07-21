package common

import (
	"testing"
)

var gcdTests = []struct {
	input    []int
	expected int
}{
	{input: []int{2, 1}, expected: 1},
	{input: []int{3, 2}, expected: 1},
	{input: []int{4, 2}, expected: 2},
	{input: []int{4, 6}, expected: 2},
	{input: []int{5, 0}, expected: 5},
	{input: []int{-20, 8}, expected: 4},
	{input: []int{24, -36}, expected: 12},
	{input: []int{-18, -12}, expected: 6},
}

func TestGcd(t *testing.T) {
	for _, tt := range gcdTests {
		actual := Gcd(tt.input[0], tt.input[1])
		if actual != tt.expected {
			t.Errorf("%v: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var lcmTests = []struct {
	input    []uint
	expected uint
}{
	{input: []uint{2, 1}, expected: 2},
	{input: []uint{3, 2}, expected: 6},
	{input: []uint{4, 2}, expected: 4},
	{input: []uint{4, 6}, expected: 12},
}

func TestLcm(t *testing.T) {
	for _, tt := range lcmTests {
		actual := Lcm(tt.input[0], tt.input[1])
		if actual != tt.expected {
			t.Errorf("%v: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}
