package common

import "fmt"

// SplitNums : 文字列を1桁の数値に分割する
func SplitNums(s string) ([]int, error) {
	if len(s) == 0 {
		return []int{}, nil
	}

	nums := make([]int, len(s))
	for i, c := range s {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("Not number: %c", c)
		}

		nums[i] = int(c - '0')
	}

	return nums, nil
}

// IsParindrome : 回文判定
func IsParindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
