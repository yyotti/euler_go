package common

import (
	"errors"
	"reflect"
	"testing"
)

var splitNumsTests = []struct {
	input    string
	expected []int
	err      error
}{
	// 正常
	{input: "", expected: []int{}},
	{input: "1", expected: []int{1}},
	{input: "12", expected: []int{1, 2}},
	{input: "012", expected: []int{0, 1, 2}},

	// 異常
	{input: "a", err: errors.New("Not number: a")},
	{input: "1b", err: errors.New("Not number: b")},
	{input: "c2", err: errors.New("Not number: c")},
	{input: "1#2", err: errors.New("Not number: #")},
	{input: "-200", err: errors.New("Not number: -")},
}

func TestSplitNums(t *testing.T) {
	for _, tt := range splitNumsTests {
		actual, err := SplitNums(tt.input)
		if err != nil {
			if tt.err == nil {
				t.Errorf("%s: Error[%s]", tt.input, err.Error())
			} else if err.Error() != tt.err.Error() {
				t.Errorf("%s: Expected error [%s] but got [%s]", tt.input, tt.err.Error(), err.Error())
			}
		} else {
			if tt.err != nil {
				t.Errorf("%s: Expected error but not", tt.input)

			} else if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s: Expected %v but got %v", tt.input, tt.expected, actual)
			}
		}
	}
}

var isParindromeTests = []struct {
	input    string
	expected bool
}{
	{input: "0", expected: true},
	{input: "1", expected: true},
	{input: "9", expected: true},
	{input: "10", expected: false},
	{input: "11", expected: true},
	{input: "12", expected: false},
	{input: "21", expected: false},
	{input: "22", expected: true},
	{input: "100", expected: false},
	{input: "101", expected: true},
	{input: "111", expected: true},
	{input: "112", expected: false},
	{input: "121", expected: true},
	{input: "1001", expected: true},
	{input: "1010", expected: false},
	{input: "2022", expected: false},
	{input: "3303", expected: false},
	{input: "4444", expected: true},
	{input: "4554", expected: true},
}

func TestIsParindrome(t *testing.T) {
	for _, tt := range isParindromeTests {
		actual := IsParindrome(tt.input)
		if actual != tt.expected {
			t.Errorf("%s: Expected %v but got %v", tt.input, tt.expected, actual)
		}
	}
}
