package main

import (
	"testing"
)

var p011Tests = []struct {
	inNums   [][]uint
	inCnt    int
	expected uint
}{
	// TEST0 {{{
	{
		inNums:   [][]uint{},
		inCnt:    0,
		expected: 0,
	},
	// }}}
	// TEST1 {{{
	{
		inNums: [][]uint{
			{1},
		},
		inCnt:    2,
		expected: 0,
	},
	// }}}
	// TEST2 {{{
	{
		inNums: [][]uint{
			{2},
		},
		inCnt:    1,
		expected: 2,
	},
	// }}}
	// TEST3 {{{
	{
		inNums: [][]uint{
			{2, 3},
			{2, 1},
		},
		inCnt:    1,
		expected: 3,
	},
	// }}}
	// TEST4 {{{
	{
		inNums: [][]uint{
			{2, 3},
			{2, 1},
		},
		inCnt:    2,
		expected: 6,
	},
	// }}}
	// TEST5 {{{
	{
		inNums: [][]uint{
			{4, 3, 2},
			{1, 3, 1},
			{1, 2, 4},
		},
		inCnt:    2,
		expected: 12,
	},
	// }}}
	// TEST6 {{{
	{
		inNums: [][]uint{
			{4, 3, 2},
			{1, 2, 1},
			{4, 3, 4},
		},
		inCnt:    3,
		expected: 48,
	},
	// }}}
}

func TestP011A(t *testing.T) {
	for i, tt := range p011Tests {
		actual := p011A(tt.inNums, tt.inCnt)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", i, tt.expected, actual)
		}
	}
}

func TestP011B(t *testing.T) {
	for i, tt := range p011Tests {
		actual := p011B(tt.inNums, tt.inCnt)
		if actual != tt.expected {
			t.Errorf("%d: Expected %d but got %d", i, tt.expected, actual)
		}
	}
}

func BenchmarkP011A(b *testing.B) {
	nums := [][]uint{
		{8, 2, 22, 97, 38, 15, 0, 40, 0, 75},
		{49, 49, 99, 40, 17, 81, 18, 57, 60, 87},
		{81, 49, 31, 73, 55, 79, 14, 29, 93, 71},
		{52, 70, 95, 23, 4, 60, 11, 42, 69, 24},
		{22, 31, 16, 71, 51, 67, 63, 89, 41, 92},
		{24, 47, 32, 60, 99, 3, 45, 2, 44, 75},
		{32, 98, 81, 28, 64, 23, 67, 10, 26, 38},
		{67, 26, 20, 68, 2, 62, 12, 20, 95, 63},
		{24, 55, 58, 5, 66, 73, 99, 26, 97, 17},
		{21, 36, 23, 9, 75, 0, 76, 44, 20, 45},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p011A(nums, 4)
	}
}

func BenchmarkP011B(b *testing.B) {
	nums := [][]uint{
		{8, 2, 22, 97, 38, 15, 0, 40, 0, 75},
		{49, 49, 99, 40, 17, 81, 18, 57, 60, 87},
		{81, 49, 31, 73, 55, 79, 14, 29, 93, 71},
		{52, 70, 95, 23, 4, 60, 11, 42, 69, 24},
		{22, 31, 16, 71, 51, 67, 63, 89, 41, 92},
		{24, 47, 32, 60, 99, 3, 45, 2, 44, 75},
		{32, 98, 81, 28, 64, 23, 67, 10, 26, 38},
		{67, 26, 20, 68, 2, 62, 12, 20, 95, 63},
		{24, 55, 58, 5, 66, 73, 99, 26, 97, 17},
		{21, 36, 23, 9, 75, 0, 76, 44, 20, 45},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p011B(nums, 4)
	}
}

func ExampleP011() {
	main()

	// Output:
	// P011A: 70600674
	// P011B: 70600674
}
