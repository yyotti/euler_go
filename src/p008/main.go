package main

import "fmt"
import "strings"
import "sync"

// The four adjacent digits in the 1000-digit number that have the greatest
// product are 9 x 9 x 8 x 9 = 5832.
//
//     73167176531330624919225119674426574742355349194934
//     96983520312774506326239578318016984801869478851843
//     85861560789112949495459501737958331952853208805511
//     12540698747158523863050715693290963295227443043557
//     66896648950445244523161731856403098711121722383113
//     62229893423380308135336276614282806444486645238749
//     30358907296290491560440772390713810515859307960866
//     70172427121883998797908792274921901699720888093776
//     65727333001053367881220235421809751254540594752243
//     52584907711670556013604839586446706324415722155397
//     53697817977846174064955149290862569321978468622482
//     83972241375657056057490261407972968652414535100474
//     82166370484403199890008895243450658541227588666881
//     16427171479924442928230863465674813919123162824586
//     17866458359124566529476545682848912883142607690042
//     24219022671055626321111109370544217506941658960408
//     07198403850962455444362981230987879927244284909188
//     84580156166097919133875499200524063689912560717606
//     05886116467109405077541002256983155200055935729725
//     71636269561882670428252483600823257530420752963450
//
// Find the thirteen adjacent digits in the 1000-digit number that have the
// greatest product. What is the value of this product?

const num = "73167176531330624919225119674426574742355349194934" +
	"96983520312774506326239578318016984801869478851843" +
	"85861560789112949495459501737958331952853208805511" +
	"12540698747158523863050715693290963295227443043557" +
	"66896648950445244523161731856403098711121722383113" +
	"62229893423380308135336276614282806444486645238749" +
	"30358907296290491560440772390713810515859307960866" +
	"70172427121883998797908792274921901699720888093776" +
	"65727333001053367881220235421809751254540594752243" +
	"52584907711670556013604839586446706324415722155397" +
	"53697817977846174064955149290862569321978468622482" +
	"83972241375657056057490261407972968652414535100474" +
	"82166370484403199890008895243450658541227588666881" +
	"16427171479924442928230863465674813919123162824586" +
	"17866458359124566529476545682848912883142607690042" +
	"24219022671055626321111109370544217506941658960408" +
	"07198403850962455444362981230987879927244284909188" +
	"84580156166097919133875499200524063689912560717606" +
	"05886116467109405077541002256983155200055935729725" +
	"71636269561882670428252483600823257530420752963450"

const count = 13

func main() {
	fmt.Printf("P008A: %d\n", p008A(count))
	fmt.Printf("P008B: %d\n", p008B(count))
	fmt.Printf("P008C: %d\n", p008C(count))
	fmt.Printf("P008D: %d\n", p008D(count))
}

func splitNums(s string) ([]uint, error) {
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

func product(nums []uint) uint {
	if len(nums) == 0 {
		return 0
	}

	p := uint(1)
	for _, n := range nums {
		p *= n
	}
	return p
}

// 数字を先頭から見ていって全パターンで積をとり、その最大値を得る
func p008A(cnt uint) uint {
	if cnt < 1 {
		return 0
	}

	max := uint(0)
	for i := 0; i <= len(num)-int(cnt); i++ {
		nums, err := splitNums(num[i : i+int(cnt)])
		if err != nil {
			panic(err)
		}

		prod := product(nums)
		if max < prod {
			max = prod
		}
	}

	return max
}

// 0が入れば数字の積は0になるので、13文字の中に0が入る場合は除外してよい。
// まず文字列を0で分割し、0を含まないそれぞれの区間ごとで指定された文字数の積
// をとる。
func p008B(cnt uint) uint {
	if cnt < 1 {
		return 0
	}

	parts := strings.Split(num, "0")
	max := uint(0)
	for _, part := range parts {
		if len(part) < int(cnt) {
			continue
		}

		for i := 0; i <= len(part)-int(cnt); i++ {
			nums, err := splitNums(part[i : i+int(cnt)])
			if err != nil {
				panic(err)
			}

			prod := product(nums)
			if max < prod {
				max = prod
			}
		}
	}

	return max
}

// ロジックはAのまま、非同期にやる
// ->やってる処理が軽すぎて、かえって遅くなった
func p008C(cnt uint) uint {
	if cnt < 1 {
		return 0
	}

	max := uint(0)

	ch := make(chan uint, len(num)-int(cnt)+1)
	done := make(chan struct{})
	defer close(done)

	go func(ch chan<- uint) {
		w := &sync.WaitGroup{}
		w.Add(len(num) - int(cnt) + 1)
		for i := 0; i <= len(num)-int(cnt); i++ {
			go func(i int) {
				defer w.Done()

				nums, err := splitNums(num[i : i+int(cnt)])
				if err != nil {
					panic(err)
				}

				ch <- product(nums)
			}(i)
		}
		w.Wait()
		close(ch)
	}(ch)

	go func(ch <-chan uint, done chan<- struct{}) {
		for n := range ch {
			if max < n {
				max = n
			}
		}
		done <- struct{}{}
	}(ch, done)

	<-done

	return max
}

// ロジックはBのまま、非同期にやる
// ->Cよりは速いけど、結局Bの方が速い
// ->ちなみにsync.Mutexを使う方がchannelより速い
func p008D(cnt uint) uint {
	if cnt < 1 {
		return 0
	}

	max := uint(0)

	ch := make(chan uint, len(num)-int(cnt)+1)
	done := make(chan struct{})
	defer close(done)

	go func(ch chan<- uint) {
		parts := strings.Split(num, "0")
		w := &sync.WaitGroup{}
		for _, part := range parts {
			if len(part) < int(cnt) {
				continue
			}

			w.Add(1)
			go func(part string) {
				defer w.Done()

				wg := &sync.WaitGroup{}
				wg.Add(len(part) - int(cnt) + 1)
				for i := 0; i <= len(part)-int(cnt); i++ {
					go func(i int) {
						defer wg.Done()

						nums, err := splitNums(part[i : i+int(cnt)])
						if err != nil {
							panic(err)
						}

						ch <- product(nums)
					}(i)
				}
				wg.Wait()
			}(part)
		}
		w.Wait()
		close(ch)
	}(ch)

	go func(ch <-chan uint, done chan<- struct{}) {
		for n := range ch {
			if max < n {
				max = n
			}
		}
		done <- struct{}{}
	}(ch, done)

	<-done

	return max
}
