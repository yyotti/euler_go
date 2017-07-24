package main

import (
	"testing"
)

var p018Tests = []struct {
	input    [][]int
	expected int
}{
	// TEST0 {{{
	{
		input:    [][]int{},
		expected: 0,
	},
	// }}}
	// TEST1 {{{
	{
		input: [][]int{
			{},
		},
		expected: -1,
	},
	// }}}
	// TEST2 {{{
	{
		input: [][]int{
			{1},
			{2},
		},
		expected: -1,
	},
	// }}}
	// TEST3 {{{
	{
		input: [][]int{
			{1},
			{2, 3},
		},
		expected: 4,
	},
	// }}}
	// TEST4 {{{
	{
		input: [][]int{
			{1},
			{2, 3},
			{2, 3, 1},
		},
		expected: 7,
	},
	// }}}
	// TEST5 {{{
	{
		input: [][]int{
			{1},
			{2, 3},
			{4},
		},
		expected: -1,
	},
	// }}}
	// TEST6 {{{
	{
		input: [][]int{
			{3},
			{7, 4},
			{2, 4, 6},
			{8, 5, 9, 3},
		},
		expected: 23,
	},
	// }}}
}

func TestP018A(t *testing.T) {
	for i, tt := range p018Tests {
		actual := p018A(tt.input)
		if actual != tt.expected {
			t.Errorf("TEST %d: Expected %d but got %d", i, tt.expected, actual)
		}
	}
}

func TestP018B(t *testing.T) {
	for i, tt := range p018Tests {
		actual := p018B(tt.input)
		if actual != tt.expected {
			t.Errorf("TEST %d: Expected %d but got %d", i, tt.expected, actual)
		}
	}
}

func TestP018C(t *testing.T) {
	for i, tt := range p018Tests {
		actual := p018C(tt.input)
		if actual != tt.expected {
			t.Errorf("TEST %d: Expected %d but got %d", i, tt.expected, actual)
		}
	}
}

var p018Bench = [][]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
}

func BenchmarkP018A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p018A(p018Bench)
	}
}

func BenchmarkP018B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p018B(p018Bench)
	}
}

func BenchmarkP018C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p018C(p018Bench)
	}
}

var p018BSubTests = []struct {
	inNums   [][]int
	inI      int
	inJ      int
	expected int
}{
	// TEST0 {{{
	{
		inNums:   [][]int{},
		inI:      0,
		inJ:      0,
		expected: -1,
	},
	// }}}
	// TEST1 {{{
	{
		inNums:   [][]int{{1}},
		inI:      -1,
		inJ:      0,
		expected: -1,
	},
	// }}}
	// TEST2 {{{
	{
		inNums:   [][]int{{1}},
		inI:      0,
		inJ:      -1,
		expected: -1,
	},
	// }}}
	// TEST3 {{{
	{
		inNums:   [][]int{{1}},
		inI:      0,
		inJ:      1,
		expected: -1,
	},
	// }}}
	// TEST4 {{{
	{
		inNums:   [][]int{{1}},
		inI:      0,
		inJ:      0,
		expected: 1,
	},
	// }}}
	// TEST5 {{{
	{
		inNums:   [][]int{{}},
		inI:      0,
		inJ:      0,
		expected: -1,
	},
	// }}}
	// TEST6 {{{
	{
		inNums: [][]int{
			{3},
			{7, 4},
		},
		inI:      1,
		inJ:      0,
		expected: 7,
	},
	// }}}
	// TEST7 {{{
	{
		inNums: [][]int{
			{3},
			{7, 4},
		},
		inI:      1,
		inJ:      1,
		expected: 4,
	},
	// }}}
	// TEST8 {{{
	{
		inNums: [][]int{
			{3},
			{7, 4},
		},
		inI:      1,
		inJ:      2,
		expected: -1,
	},
	// }}}
	// TEST9 {{{
	{
		inNums: [][]int{
			{3},
			{7, 4},
		},
		inI:      0,
		inJ:      0,
		expected: 10,
	},
	// }}}
	// TEST10 {{{
	{
		inNums: [][]int{
			{3},
			{7, 4},
			{2, 4, 6},
		},
		inI:      1,
		inJ:      0,
		expected: 11,
	},
	// }}}
	// TEST11 {{{
	{
		inNums: [][]int{
			{3},
			{7, 4},
			{2, 4, 6},
		},
		inI:      1,
		inJ:      1,
		expected: 10,
	},
	// }}}
	// TEST12 {{{
	{
		inNums: [][]int{
			{3},
			{7, 4},
			{2, 4, 6},
		},
		inI:      0,
		inJ:      0,
		expected: 14,
	},
	// }}}
}

func TestP018BSub(t *testing.T) {
	for i, tt := range p018BSubTests {
		actual := p018BSub(tt.inNums, tt.inI, tt.inJ)
		if actual != tt.expected {
			t.Errorf("TEST %d: Expected %d but got %d", i, tt.expected, actual)
		}
	}
}

func Example018() {
	main()

	// Output:
	// P018A: 1074
	// P018B: 1074
	// P018C: 1074
}
