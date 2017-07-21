package common

import (
	"errors"
	"reflect"
	"testing"
)

var splitNumsTests = []struct {
	input    string
	expected []uint
	err      error
}{
	// 正常
	{input: "", expected: []uint{}},
	{input: "1", expected: []uint{1}},
	{input: "12", expected: []uint{1, 2}},
	{input: "012", expected: []uint{0, 1, 2}},

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
