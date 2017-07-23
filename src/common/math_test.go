package common

import (
	"math/big"
	"reflect"
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

var permutationATests = []struct {
	inN      uint
	inR      uint
	expected uint64
}{
	{inN: 0, inR: 0, expected: 1},
	{inN: 1, inR: 0, expected: 1},
	{inN: 1, inR: 1, expected: 1},
	{inN: 1, inR: 2, expected: 1},
	{inN: 2, inR: 0, expected: 1},
	{inN: 2, inR: 1, expected: 2},
	{inN: 2, inR: 2, expected: 2},
	{inN: 2, inR: 3, expected: 2},
	{inN: 3, inR: 0, expected: 1},
	{inN: 3, inR: 1, expected: 3},
	{inN: 3, inR: 2, expected: 6},
	{inN: 3, inR: 3, expected: 6},
	{inN: 3, inR: 4, expected: 6},
	{inN: 4, inR: 0, expected: 1},
	{inN: 4, inR: 1, expected: 4},
	{inN: 4, inR: 2, expected: 12},
	{inN: 4, inR: 3, expected: 24},
	{inN: 4, inR: 4, expected: 24},
	{inN: 4, inR: 5, expected: 24},
}

func TestPermutationA(t *testing.T) {
	for _, tt := range permutationATests {
		actual := PermutationA(tt.inN, tt.inR)
		if actual != tt.expected {
			t.Errorf("P(%d,%d): Expected %d but got %d", tt.inN, tt.inR, tt.expected, actual)
		}
	}
}

var permutationBTests = []struct {
	inN      uint
	inR      uint
	expected *big.Int
}{
	{inN: 0, inR: 0, expected: big.NewInt(1)},
	{inN: 1, inR: 0, expected: big.NewInt(1)},
	{inN: 1, inR: 1, expected: big.NewInt(1)},
	{inN: 1, inR: 2, expected: big.NewInt(1)},
	{inN: 2, inR: 0, expected: big.NewInt(1)},
	{inN: 2, inR: 1, expected: big.NewInt(2)},
	{inN: 2, inR: 2, expected: big.NewInt(2)},
	{inN: 2, inR: 3, expected: big.NewInt(2)},
	{inN: 3, inR: 0, expected: big.NewInt(1)},
	{inN: 3, inR: 1, expected: big.NewInt(3)},
	{inN: 3, inR: 2, expected: big.NewInt(6)},
	{inN: 3, inR: 3, expected: big.NewInt(6)},
	{inN: 3, inR: 4, expected: big.NewInt(6)},
	{inN: 4, inR: 0, expected: big.NewInt(1)},
	{inN: 4, inR: 1, expected: big.NewInt(4)},
	{inN: 4, inR: 2, expected: big.NewInt(12)},
	{inN: 4, inR: 3, expected: big.NewInt(24)},
	{inN: 4, inR: 4, expected: big.NewInt(24)},
	{inN: 4, inR: 5, expected: big.NewInt(24)},
}

func TestPermutationB(t *testing.T) {
	for _, tt := range permutationBTests {
		actual := PermutationB(tt.inN, tt.inR)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("P(%d,%d): Expected %d but got %d", tt.inN, tt.inR, tt.expected.Int64(), actual.Int64())
		}
	}
}

func BenchmarkPermutationA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PermutationA(20, 10)
	}
}

func BenchmarkPermutationB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PermutationB(20, 10)
	}
}

var combinationTests = []struct {
	inN      uint
	inR      uint
	expected uint64
}{
	{inN: 0, inR: 0, expected: 1},
	{inN: 1, inR: 0, expected: 1},
	{inN: 1, inR: 1, expected: 1},
	{inN: 1, inR: 2, expected: 1},
	{inN: 2, inR: 0, expected: 1},
	{inN: 2, inR: 1, expected: 2},
	{inN: 2, inR: 2, expected: 1},
	{inN: 2, inR: 3, expected: 1},
	{inN: 3, inR: 0, expected: 1},
	{inN: 3, inR: 1, expected: 3},
	{inN: 3, inR: 2, expected: 3},
	{inN: 3, inR: 3, expected: 1},
	{inN: 3, inR: 4, expected: 1},
	{inN: 4, inR: 0, expected: 1},
	{inN: 4, inR: 1, expected: 4},
	{inN: 4, inR: 2, expected: 6},
	{inN: 4, inR: 3, expected: 4},
	{inN: 4, inR: 4, expected: 1},
	{inN: 4, inR: 5, expected: 1},
	{inN: 5, inR: 0, expected: 1},
	{inN: 5, inR: 1, expected: 5},
	{inN: 5, inR: 2, expected: 10},
	{inN: 5, inR: 3, expected: 10},
	{inN: 5, inR: 4, expected: 5},
	{inN: 5, inR: 5, expected: 1},
	{inN: 5, inR: 6, expected: 1},
}

func TestCombinationA(t *testing.T) {
	for _, tt := range combinationTests {
		actual := CombinationA(tt.inN, tt.inR)
		if actual != tt.expected {
			t.Errorf("P(%d,%d): Expected %d but got %d", tt.inN, tt.inR, tt.expected, actual)
		}
	}
}

func TestCombinationB(t *testing.T) {
	for _, tt := range combinationTests {
		actual := CombinationB(tt.inN, tt.inR)
		if actual != tt.expected {
			t.Errorf("P(%d,%d): Expected %d but got %d", tt.inN, tt.inR, tt.expected, actual)
		}
	}
}

func TestCombinationC(t *testing.T) {
	for _, tt := range combinationTests {
		actual := CombinationC(tt.inN, tt.inR)
		if actual != tt.expected {
			t.Errorf("P(%d,%d): Expected %d but got %d", tt.inN, tt.inR, tt.expected, actual)
		}
	}
}

func TestCombinationD(t *testing.T) {
	for _, tt := range combinationTests {
		actual := CombinationD(tt.inN, tt.inR)
		if actual != tt.expected {
			t.Errorf("P(%d,%d): Expected %d but got %d", tt.inN, tt.inR, tt.expected, actual)
		}
	}
}

func BenchmarkCombinationA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombinationA(20, 10)
	}
}

func BenchmarkCombinationB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombinationB(20, 10)
	}
}

func BenchmarkCombinationC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombinationC(20, 10)
	}
}

func BenchmarkCombinationD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombinationD(20, 10)
	}
}
