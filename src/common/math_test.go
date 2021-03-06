package common

import (
	"math/big"
	"reflect"
	"sort"
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
	input    []int
	expected int
}{
	{input: []int{2, 1}, expected: 2},
	{input: []int{-2, 1}, expected: 2},
	{input: []int{2, -1}, expected: 2},
	{input: []int{-2, -1}, expected: 2},
	{input: []int{3, 2}, expected: 6},
	{input: []int{4, 2}, expected: 4},
	{input: []int{4, 6}, expected: 12},
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
	inN      int
	inR      int
	expected int64
}{
	{inN: -1, inR: -2, expected: 0},
	{inN: -1, inR: -1, expected: 0},
	{inN: 1, inR: -1, expected: 0},
	{inN: -1, inR: 1, expected: 0},
	{inN: 0, inR: 0, expected: 1},
	{inN: 1, inR: 0, expected: 1},
	{inN: 1, inR: 1, expected: 1},
	{inN: 1, inR: 2, expected: 0},
	{inN: 2, inR: 0, expected: 1},
	{inN: 2, inR: 1, expected: 2},
	{inN: 2, inR: 2, expected: 2},
	{inN: 2, inR: 3, expected: 0},
	{inN: 3, inR: 0, expected: 1},
	{inN: 3, inR: 1, expected: 3},
	{inN: 3, inR: 2, expected: 6},
	{inN: 3, inR: 3, expected: 6},
	{inN: 3, inR: 4, expected: 0},
	{inN: 4, inR: 0, expected: 1},
	{inN: 4, inR: 1, expected: 4},
	{inN: 4, inR: 2, expected: 12},
	{inN: 4, inR: 3, expected: 24},
	{inN: 4, inR: 4, expected: 24},
	{inN: 4, inR: 5, expected: 0},
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
	inN      int
	inR      int
	expected *big.Int
}{
	{inN: -1, inR: -2, expected: big.NewInt(0)},
	{inN: -1, inR: -1, expected: big.NewInt(0)},
	{inN: 1, inR: -1, expected: big.NewInt(0)},
	{inN: -1, inR: 1, expected: big.NewInt(0)},
	{inN: 0, inR: 0, expected: big.NewInt(1)},
	{inN: 1, inR: 0, expected: big.NewInt(1)},
	{inN: 1, inR: 1, expected: big.NewInt(1)},
	{inN: 1, inR: 2, expected: big.NewInt(0)},
	{inN: 2, inR: 0, expected: big.NewInt(1)},
	{inN: 2, inR: 1, expected: big.NewInt(2)},
	{inN: 2, inR: 2, expected: big.NewInt(2)},
	{inN: 2, inR: 3, expected: big.NewInt(0)},
	{inN: 3, inR: 0, expected: big.NewInt(1)},
	{inN: 3, inR: 1, expected: big.NewInt(3)},
	{inN: 3, inR: 2, expected: big.NewInt(6)},
	{inN: 3, inR: 3, expected: big.NewInt(6)},
	{inN: 3, inR: 4, expected: big.NewInt(0)},
	{inN: 4, inR: 0, expected: big.NewInt(1)},
	{inN: 4, inR: 1, expected: big.NewInt(4)},
	{inN: 4, inR: 2, expected: big.NewInt(12)},
	{inN: 4, inR: 3, expected: big.NewInt(24)},
	{inN: 4, inR: 4, expected: big.NewInt(24)},
	{inN: 4, inR: 5, expected: big.NewInt(0)},
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
	inN      int
	inR      int
	expected int64
}{
	{inN: -1, inR: -1, expected: 0},
	{inN: 1, inR: -1, expected: 0},
	{inN: -1, inR: 1, expected: 0},
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

var addTests = []struct {
	inA      string
	inB      string
	expected string
}{
	{inA: "0", inB: "0", expected: "0"},
	{inA: "1", inB: "0", expected: "1"},
	{inA: "0", inB: "2", expected: "2"},
	{inA: "1", inB: "2", expected: "3"},
	{inA: "10", inB: "1", expected: "11"},
	{inA: "2", inB: "20", expected: "22"},
	{inA: "12", inB: "23", expected: "35"},
	{inA: "1234", inB: "5678", expected: "6912"},
	{inA: "2", inB: "4998", expected: "5000"},
	{inA: "99", inB: "9", expected: "108"},
}

func TestAdd(t *testing.T) {
	for _, tt := range addTests {
		actual := Add(tt.inA, tt.inB)
		if actual != tt.expected {
			t.Errorf("%s+%s: Expected %s but got %s", tt.inA, tt.inB, tt.expected, actual)
		}
	}
}

var mulTests = []struct {
	inA      string
	inB      string
	expected string
}{
	{inA: "0", inB: "0", expected: "0"},
	{inA: "1", inB: "0", expected: "0"},
	{inA: "0", inB: "2", expected: "0"},
	{inA: "1", inB: "2", expected: "2"},
	{inA: "10", inB: "1", expected: "10"},
	{inA: "2", inB: "20", expected: "40"},
	{inA: "12", inB: "23", expected: "276"},
	{inA: "1", inB: "200", expected: "200"},
}

func TestMul(t *testing.T) {
	for _, tt := range mulTests {
		actual := Mul(tt.inA, tt.inB)
		if actual != tt.expected {
			t.Errorf("%s+%s: Expected %s but got %s", tt.inA, tt.inB, tt.expected, actual)
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
		actual := Divisors(tt.input)
		sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divisors(3628800) // 1 x 2 x ... x 10
	}
}

var permutationNTests = []struct {
	inNums   []int
	inN      int
	expected []int
}{
	{inNums: []int{}, inN: 0, expected: []int{}},
	{inNums: []int{}, inN: 1, expected: []int{}},
	{inNums: []int{0}, inN: 0, expected: []int{}},
	{inNums: []int{0}, inN: 1, expected: []int{0}},
	{inNums: []int{0}, inN: 2, expected: []int{}},
	{inNums: []int{1, 2}, inN: 1, expected: []int{1, 2}},
	{inNums: []int{1, 2}, inN: 2, expected: []int{2, 1}},
	{inNums: []int{0, 1, 2}, inN: 0, expected: []int{}},
	{inNums: []int{0, 1, 2}, inN: 1, expected: []int{0, 1, 2}},
	{inNums: []int{0, 1, 2}, inN: 2, expected: []int{0, 2, 1}},
	{inNums: []int{0, 1, 2}, inN: 3, expected: []int{1, 0, 2}},
	{inNums: []int{0, 1, 2}, inN: 4, expected: []int{1, 2, 0}},
	{inNums: []int{0, 1, 2}, inN: 5, expected: []int{2, 0, 1}},
	{inNums: []int{0, 1, 2}, inN: 6, expected: []int{2, 1, 0}},
	{inNums: []int{0, 1, 2}, inN: 7, expected: []int{}},
}

func TestPermutationN(t *testing.T) {
	for _, tt := range permutationNTests {
		actual := NthPermutation(tt.inNums, tt.inN)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%v,%d: Expected %v but got %v", tt.inNums, tt.inN, tt.expected, actual)
		}
	}
}

func BenchmarkPermutationN(b *testing.B) {
	nums := []int{0, 1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NthPermutation(nums, 120)
	}
}

var factoradicTests = []struct {
	input    int
	expected []int
}{
	{input: -1, expected: nil},
	{input: 0, expected: []int{}},
	{input: 1, expected: []int{1}},
	{input: 2, expected: []int{1, 0}},
	{input: 3, expected: []int{1, 1}},
	{input: 4, expected: []int{2, 0}},
	{input: 5, expected: []int{2, 1}},
	{input: 6, expected: []int{1, 0, 0}},
	{input: 7, expected: []int{1, 0, 1}},
	{input: 8, expected: []int{1, 1, 0}},
	{input: 9, expected: []int{1, 1, 1}},
	{input: 10, expected: []int{1, 2, 0}},
	{input: 11, expected: []int{1, 2, 1}},
	{input: 12, expected: []int{2, 0, 0}},
	{input: 13, expected: []int{2, 0, 1}},
	{input: 14, expected: []int{2, 1, 0}},
	{input: 15, expected: []int{2, 1, 1}},
	{input: 16, expected: []int{2, 2, 0}},
	{input: 17, expected: []int{2, 2, 1}},
	{input: 18, expected: []int{3, 0, 0}},
	{input: 19, expected: []int{3, 0, 1}},
	{input: 20, expected: []int{3, 1, 0}},
	{input: 21, expected: []int{3, 1, 1}},
	{input: 22, expected: []int{3, 2, 0}},
	{input: 23, expected: []int{3, 2, 1}},
}

func TestFactoradic(t *testing.T) {
	for _, tt := range factoradicTests {
		actual := factoradic(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%d: Expected %d but got %d", tt.input, tt.expected, actual)
		}
	}
}

var isParindromeNumTests = []struct {
	input    int
	expected bool
}{
	{input: 0, expected: true},
	{input: 1, expected: true},
	{input: 9, expected: true},
	{input: 10, expected: false},
	{input: 11, expected: true},
	{input: 12, expected: false},
	{input: 21, expected: false},
	{input: 22, expected: true},
	{input: 100, expected: false},
	{input: 101, expected: true},
	{input: 111, expected: true},
	{input: 112, expected: false},
	{input: 121, expected: true},
	{input: 1001, expected: true},
	{input: 1010, expected: false},
	{input: 2022, expected: false},
	{input: 3303, expected: false},
	{input: 4444, expected: true},
	{input: 4554, expected: true},
}

func TestIsParindromeNum(t *testing.T) {
	for _, tt := range isParindromeNumTests {
		actual := IsParindromeNum(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}
