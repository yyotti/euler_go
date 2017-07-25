package main

import "reflect"
import "testing"

var permutationsTests = []struct {
	input    []int
	expected [][]int
}{
	{input: []int{}, expected: [][]int{}},
	{input: []int{0}, expected: [][]int{{0}}},
	{input: []int{0, 1}, expected: [][]int{{0, 1}, {1, 0}}},
	{
		input: []int{0, 1, 2},
		expected: [][]int{
			{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0},
		},
	},
}

func TestPermutations(t *testing.T) {
	for _, tt := range permutationsTests {
		actual := permutations(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%v: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkPermutations(b *testing.B) {
	nums := []int{0, 1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		permutations(nums)
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

var p024Tests = []struct {
	inNums   []int
	inN      int
	expected int
}{
	{inNums: []int{}, inN: 0, expected: -1},
	{inNums: []int{}, inN: 1, expected: -1},
	{inNums: []int{0}, inN: 0, expected: -1},
	{inNums: []int{0}, inN: 1, expected: 0},
	{inNums: []int{0}, inN: 2, expected: -1},
	{inNums: []int{1, 2}, inN: 1, expected: 12},
	{inNums: []int{1, 2}, inN: 2, expected: 21},
	{inNums: []int{0, 1, 2}, inN: 0, expected: -1},
	{inNums: []int{0, 1, 2}, inN: 1, expected: 12},
	{inNums: []int{0, 1, 2}, inN: 2, expected: 21},
	{inNums: []int{0, 1, 2}, inN: 3, expected: 102},
	{inNums: []int{0, 1, 2}, inN: 4, expected: 120},
	{inNums: []int{0, 1, 2}, inN: 5, expected: 201},
	{inNums: []int{0, 1, 2}, inN: 6, expected: 210},
	{inNums: []int{0, 1, 2}, inN: 7, expected: -1},
}

func TestP024A(t *testing.T) {
	for _, tt := range p024Tests {
		actual := p024A(tt.inNums, tt.inN)
		if actual != tt.expected {
			t.Errorf("(%v,%d): Expected %d but got %d", tt.inNums, tt.inN, tt.expected, actual)
		}
	}
}

func TestP024B(t *testing.T) {
	for _, tt := range p024Tests {
		actual := p024B(tt.inNums, tt.inN)
		if actual != tt.expected {
			t.Errorf("(%v,%d): Expected %d but got %d", tt.inNums, tt.inN, tt.expected, actual)
		}
	}
}

func BenchmarkP024A(b *testing.B) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p024A(nums, 30000)
	}
}

func BenchmarkP024B(b *testing.B) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p024B(nums, 30000)
	}
}

func ExampleP024() {
	main()

	// Output:
	// P024A([0-7],30000): 57364210
	// P024B: 2783915460
}
