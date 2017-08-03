package main

import "fmt"

// In England the currency is made up of pound, £, and pence, p, and there are
// eight coins in general circulation:
//
//     1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
//
// It is possible to make £2 in the following way:
//
//     1x£1 + 1x50p + 2x20p + 1x5p + 1x2p + 3x1p
//
// How many different ways can £2 be made using any number of coins?

var coins = []int{200, 100, 50, 20, 10, 5, 2, 1}

const price = 200

func main() {
	fmt.Printf("P031A: %d\n", p031A(price, coins))
}

func p031A(p int, coins []int) int {
	if p < 1 || len(coins) == 0 {
		return 0
	}

	cnt := 0
	for i, c := range coins {
		for j := p / c; j > 0; j-- {
			if p == c*j {
				cnt++
				continue
			}

			cnt += p031A(p-c*j, coins[i+1:])
		}
	}

	return cnt
}
