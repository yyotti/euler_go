package common

import "reflect"
import "testing"

var digitCountTests = []struct {
	input    uint64
	expected int
}{
	{input: 0, expected: 1},
	{input: 1, expected: 1},
	{input: 2, expected: 1},
	{input: 10, expected: 2},
	{input: 11, expected: 2},
	{input: 23, expected: 2},
	{input: 200, expected: 3},
	{input: 201, expected: 3},
	{input: 231, expected: 3},
}

func TestDigitCount(t *testing.T) {
	for _, tt := range digitCountTests {
		actual := DigitCount(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Exptected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var digitsTests = []struct {
	input    uint64
	expected []int
}{
	{input: 0, expected: []int{0}},
	{input: 1, expected: []int{1}},
	{input: 2, expected: []int{2}},
	{input: 10, expected: []int{1, 0}},
	{input: 11, expected: []int{1, 1}},
	{input: 23, expected: []int{2, 3}},
	{input: 200, expected: []int{2, 0, 0}},
	{input: 201, expected: []int{2, 0, 1}},
	{input: 231, expected: []int{2, 3, 1}},
}

func TestDigits(t *testing.T) {
	for _, tt := range digitsTests {
		actual := Digits(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Exptected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var isPandigitalTests = []struct {
	inN      int
	inM      int
	expected bool
}{
	{inN: 12, inM: -1, expected: false},
	{inN: -1, inM: 0, expected: false},
	{inN: -2, inM: 1, expected: false},
	{inN: 0, inM: 1, expected: false},
	{inN: 1, inM: 1, expected: true},
	{inN: 2, inM: 1, expected: false},
	{inN: 1, inM: 2, expected: false},
	{inN: 11, inM: 1, expected: false},
	{inN: 12, inM: 1, expected: false},
	{inN: 12, inM: 2, expected: true},
	{inN: 12, inM: 3, expected: false},
	{inN: 13, inM: 2, expected: false},
	{inN: 13, inM: 3, expected: false},
	{inN: 21, inM: 2, expected: true},
	{inN: 120, inM: 3, expected: false},
	{inN: 123, inM: 2, expected: false},
	{inN: 123, inM: 3, expected: true},
	{inN: 132, inM: 3, expected: true},
}

func TestIsPandigital(t *testing.T) {
	for _, tt := range isPandigitalTests {
		actual := IsPandigital(tt.inN, tt.inM)
		if actual != tt.expected {
			t.Errorf("%d,%d: Exptected %v but got %v", tt.inN, tt.inM, tt.expected, actual)
		}
	}
}
