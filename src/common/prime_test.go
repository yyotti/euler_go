package common

import (
	"reflect"
	"testing"
)

var primeFactorsTests = []struct {
	input    uint
	expected map[uint]uint
}{
	{input: 0, expected: map[uint]uint{}},
	{input: 1, expected: map[uint]uint{}},
	{input: 2, expected: map[uint]uint{2: 1}},
	{input: 3, expected: map[uint]uint{3: 1}},
	{input: 4, expected: map[uint]uint{2: 2}},
	{input: 5, expected: map[uint]uint{5: 1}},
	{input: 6, expected: map[uint]uint{2: 1, 3: 1}},
	{input: 7, expected: map[uint]uint{7: 1}},
	{input: 8, expected: map[uint]uint{2: 3}},
	{input: 9, expected: map[uint]uint{3: 2}},
	{input: 13195, expected: map[uint]uint{5: 1, 7: 1, 13: 1, 29: 1}},
}

func TestPrimeFactors(t *testing.T) {
	for _, tt := range primeFactorsTests {
		actual := PrimeFactors(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}
