package main

import "testing"

const p019Expected = 171

var isLeapYearTests = []struct {
	input    int
	expected bool
}{
	{input: 1900, expected: false},
	{input: 1901, expected: false},
	{input: 1902, expected: false},
	{input: 1903, expected: false},
	{input: 1904, expected: true},
	{input: 1905, expected: false},
	{input: 1916, expected: true},
	{input: 2000, expected: true},
	{input: 2001, expected: false},
	{input: 2002, expected: false},
	{input: 2003, expected: false},
	{input: 2004, expected: true},
	{input: 2020, expected: true},
	{input: 2100, expected: false},
}

func TestIsLeapYear(t *testing.T) {
	for _, tt := range isLeapYearTests {
		actual := isLeapYear(tt.input)
		if actual != tt.expected {
			t.Errorf("%d: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}

var datesTests = []struct {
	inY      int
	inM      int
	expected int
}{
	// 1900 {{{
	{inY: 1900, inM: 1, expected: 31},
	{inY: 1900, inM: 2, expected: 28},
	{inY: 1900, inM: 3, expected: 31},
	{inY: 1900, inM: 4, expected: 30},
	{inY: 1900, inM: 5, expected: 31},
	{inY: 1900, inM: 6, expected: 30},
	{inY: 1900, inM: 7, expected: 31},
	{inY: 1900, inM: 8, expected: 31},
	{inY: 1900, inM: 9, expected: 30},
	{inY: 1900, inM: 10, expected: 31},
	{inY: 1900, inM: 11, expected: 30},
	{inY: 1900, inM: 12, expected: 31},
	// }}}
	// 1901 {{{
	{inY: 1901, inM: 1, expected: 31},
	{inY: 1901, inM: 2, expected: 28},
	{inY: 1901, inM: 3, expected: 31},
	{inY: 1901, inM: 4, expected: 30},
	{inY: 1901, inM: 5, expected: 31},
	{inY: 1901, inM: 6, expected: 30},
	{inY: 1901, inM: 7, expected: 31},
	{inY: 1901, inM: 8, expected: 31},
	{inY: 1901, inM: 9, expected: 30},
	{inY: 1901, inM: 10, expected: 31},
	{inY: 1901, inM: 11, expected: 30},
	{inY: 1901, inM: 12, expected: 31},
	// }}}
	// 1902 {{{
	{inY: 1902, inM: 1, expected: 31},
	{inY: 1902, inM: 2, expected: 28},
	{inY: 1902, inM: 3, expected: 31},
	{inY: 1902, inM: 4, expected: 30},
	{inY: 1902, inM: 5, expected: 31},
	{inY: 1902, inM: 6, expected: 30},
	{inY: 1902, inM: 7, expected: 31},
	{inY: 1902, inM: 8, expected: 31},
	{inY: 1902, inM: 9, expected: 30},
	{inY: 1902, inM: 10, expected: 31},
	{inY: 1902, inM: 11, expected: 30},
	{inY: 1902, inM: 12, expected: 31},
	// }}}
	// 1903 {{{
	{inY: 1903, inM: 1, expected: 31},
	{inY: 1903, inM: 2, expected: 28},
	{inY: 1903, inM: 3, expected: 31},
	{inY: 1903, inM: 4, expected: 30},
	{inY: 1903, inM: 5, expected: 31},
	{inY: 1903, inM: 6, expected: 30},
	{inY: 1903, inM: 7, expected: 31},
	{inY: 1903, inM: 8, expected: 31},
	{inY: 1903, inM: 9, expected: 30},
	{inY: 1903, inM: 10, expected: 31},
	{inY: 1903, inM: 11, expected: 30},
	{inY: 1903, inM: 12, expected: 31},
	// }}}
	// 1904 {{{
	{inY: 1904, inM: 1, expected: 31},
	{inY: 1904, inM: 2, expected: 29},
	{inY: 1904, inM: 3, expected: 31},
	{inY: 1904, inM: 4, expected: 30},
	{inY: 1904, inM: 5, expected: 31},
	{inY: 1904, inM: 6, expected: 30},
	{inY: 1904, inM: 7, expected: 31},
	{inY: 1904, inM: 8, expected: 31},
	{inY: 1904, inM: 9, expected: 30},
	{inY: 1904, inM: 10, expected: 31},
	{inY: 1904, inM: 11, expected: 30},
	{inY: 1904, inM: 12, expected: 31},
	// }}}
	// 1905 {{{
	{inY: 1905, inM: 1, expected: 31},
	{inY: 1905, inM: 2, expected: 28},
	{inY: 1905, inM: 3, expected: 31},
	{inY: 1905, inM: 4, expected: 30},
	{inY: 1905, inM: 5, expected: 31},
	{inY: 1905, inM: 6, expected: 30},
	{inY: 1905, inM: 7, expected: 31},
	{inY: 1905, inM: 8, expected: 31},
	{inY: 1905, inM: 9, expected: 30},
	{inY: 1905, inM: 10, expected: 31},
	{inY: 1905, inM: 11, expected: 30},
	{inY: 1905, inM: 12, expected: 31},
	// }}}
	// 1916 {{{
	{inY: 1916, inM: 1, expected: 31},
	{inY: 1916, inM: 2, expected: 29},
	{inY: 1916, inM: 3, expected: 31},
	{inY: 1916, inM: 4, expected: 30},
	{inY: 1916, inM: 5, expected: 31},
	{inY: 1916, inM: 6, expected: 30},
	{inY: 1916, inM: 7, expected: 31},
	{inY: 1916, inM: 8, expected: 31},
	{inY: 1916, inM: 9, expected: 30},
	{inY: 1916, inM: 10, expected: 31},
	{inY: 1916, inM: 11, expected: 30},
	{inY: 1916, inM: 12, expected: 31},
	// }}}
	// 2000 {{{
	{inY: 2000, inM: 1, expected: 31},
	{inY: 2000, inM: 2, expected: 29},
	{inY: 2000, inM: 3, expected: 31},
	{inY: 2000, inM: 4, expected: 30},
	{inY: 2000, inM: 5, expected: 31},
	{inY: 2000, inM: 6, expected: 30},
	{inY: 2000, inM: 7, expected: 31},
	{inY: 2000, inM: 8, expected: 31},
	{inY: 2000, inM: 9, expected: 30},
	{inY: 2000, inM: 10, expected: 31},
	{inY: 2000, inM: 11, expected: 30},
	{inY: 2000, inM: 12, expected: 31},
	// }}}
	// 2001 {{{
	{inY: 2001, inM: 1, expected: 31},
	{inY: 2001, inM: 2, expected: 28},
	{inY: 2001, inM: 3, expected: 31},
	{inY: 2001, inM: 4, expected: 30},
	{inY: 2001, inM: 5, expected: 31},
	{inY: 2001, inM: 6, expected: 30},
	{inY: 2001, inM: 7, expected: 31},
	{inY: 2001, inM: 8, expected: 31},
	{inY: 2001, inM: 9, expected: 30},
	{inY: 2001, inM: 10, expected: 31},
	{inY: 2001, inM: 11, expected: 30},
	{inY: 2001, inM: 12, expected: 31},
	// }}}
	// 2002 {{{
	{inY: 2002, inM: 1, expected: 31},
	{inY: 2002, inM: 2, expected: 28},
	{inY: 2002, inM: 3, expected: 31},
	{inY: 2002, inM: 4, expected: 30},
	{inY: 2002, inM: 5, expected: 31},
	{inY: 2002, inM: 6, expected: 30},
	{inY: 2002, inM: 7, expected: 31},
	{inY: 2002, inM: 8, expected: 31},
	{inY: 2002, inM: 9, expected: 30},
	{inY: 2002, inM: 10, expected: 31},
	{inY: 2002, inM: 11, expected: 30},
	{inY: 2002, inM: 12, expected: 31},
	// }}}
	// 2003 {{{
	{inY: 2003, inM: 1, expected: 31},
	{inY: 2003, inM: 2, expected: 28},
	{inY: 2003, inM: 3, expected: 31},
	{inY: 2003, inM: 4, expected: 30},
	{inY: 2003, inM: 5, expected: 31},
	{inY: 2003, inM: 6, expected: 30},
	{inY: 2003, inM: 7, expected: 31},
	{inY: 2003, inM: 8, expected: 31},
	{inY: 2003, inM: 9, expected: 30},
	{inY: 2003, inM: 10, expected: 31},
	{inY: 2003, inM: 11, expected: 30},
	{inY: 2003, inM: 12, expected: 31},
	// }}}
	// 2004 {{{
	{inY: 2004, inM: 1, expected: 31},
	{inY: 2004, inM: 2, expected: 29},
	{inY: 2004, inM: 3, expected: 31},
	{inY: 2004, inM: 4, expected: 30},
	{inY: 2004, inM: 5, expected: 31},
	{inY: 2004, inM: 6, expected: 30},
	{inY: 2004, inM: 7, expected: 31},
	{inY: 2004, inM: 8, expected: 31},
	{inY: 2004, inM: 9, expected: 30},
	{inY: 2004, inM: 10, expected: 31},
	{inY: 2004, inM: 11, expected: 30},
	{inY: 2004, inM: 12, expected: 31},
	// }}}
	// 2020 {{{
	{inY: 2020, inM: 1, expected: 31},
	{inY: 2020, inM: 2, expected: 29},
	{inY: 2020, inM: 3, expected: 31},
	{inY: 2020, inM: 4, expected: 30},
	{inY: 2020, inM: 5, expected: 31},
	{inY: 2020, inM: 6, expected: 30},
	{inY: 2020, inM: 7, expected: 31},
	{inY: 2020, inM: 8, expected: 31},
	{inY: 2020, inM: 9, expected: 30},
	{inY: 2020, inM: 10, expected: 31},
	{inY: 2020, inM: 11, expected: 30},
	{inY: 2020, inM: 12, expected: 31},
	// }}}
	// 2100 {{{
	{inY: 2100, inM: 1, expected: 31},
	{inY: 2100, inM: 2, expected: 28},
	{inY: 2100, inM: 3, expected: 31},
	{inY: 2100, inM: 4, expected: 30},
	{inY: 2100, inM: 5, expected: 31},
	{inY: 2100, inM: 6, expected: 30},
	{inY: 2100, inM: 7, expected: 31},
	{inY: 2100, inM: 8, expected: 31},
	{inY: 2100, inM: 9, expected: 30},
	{inY: 2100, inM: 10, expected: 31},
	{inY: 2100, inM: 11, expected: 30},
	{inY: 2100, inM: 12, expected: 31},
	// }}}
}

func TestDates(t *testing.T) {
	for _, tt := range datesTests {
		actual := dates(tt.inY, tt.inM)
		if actual != tt.expected {
			t.Errorf("%d/%d: Expected %d but got %d", tt.inY, tt.inM, tt.expected, actual)
		}
	}
}

var nextTests = []struct {
	inY       int
	inM       int
	expectedY int
	expectedM int
}{
	{inY: 1900, inM: 1, expectedY: 1900, expectedM: 2},
	{inY: 1900, inM: 2, expectedY: 1900, expectedM: 3},
	{inY: 1900, inM: 3, expectedY: 1900, expectedM: 4},
	{inY: 1900, inM: 4, expectedY: 1900, expectedM: 5},
	{inY: 1900, inM: 5, expectedY: 1900, expectedM: 6},
	{inY: 1900, inM: 6, expectedY: 1900, expectedM: 7},
	{inY: 1900, inM: 7, expectedY: 1900, expectedM: 8},
	{inY: 1900, inM: 8, expectedY: 1900, expectedM: 9},
	{inY: 1900, inM: 9, expectedY: 1900, expectedM: 10},
	{inY: 1900, inM: 10, expectedY: 1900, expectedM: 11},
	{inY: 1900, inM: 11, expectedY: 1900, expectedM: 12},
	{inY: 1900, inM: 12, expectedY: 1901, expectedM: 1},
	{inY: 1901, inM: 1, expectedY: 1901, expectedM: 2},
}

func TestNext(t *testing.T) {
	for _, tt := range nextTests {
		actualY, actualM := next(tt.inY, tt.inM)
		if actualY != tt.expectedY || actualM != tt.expectedM {
			t.Errorf("%d/%d: Expected %d/%d but got %d/%d", tt.inY, tt.inM, tt.expectedY, tt.expectedM, actualY, actualM)
		}
	}
}

func TestP019A(t *testing.T) {
	actual := p019A()
	if actual != p019Expected {
		t.Errorf("Expected %d but got %d", p019Expected, actual)
	}
}

func TestP019B(t *testing.T) {
	actual := p019B()
	if actual != p019Expected {
		t.Errorf("Expected %d but got %d", p019Expected, actual)
	}
}

func BenchmarkP019A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p019A()
	}
}

func BenchmarkP019B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p019B()
	}
}

func ExampleP019() {
	main()

	// Output:
	// P019A: 171
}
