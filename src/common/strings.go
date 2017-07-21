package common

import "fmt"

// SplitNums : 文字列を1桁の数値に分割する
func SplitNums(s string) ([]uint, error) {
	if len(s) == 0 {
		return []uint{}, nil
	}

	nums := make([]uint, len(s))
	for i, c := range s {
		if c < '0' || c > '9' {
			return nil, fmt.Errorf("Not number: %c", c)
		}

		nums[i] = uint(c - '0')
	}

	return nums, nil
}
