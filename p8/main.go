package main

import (
	"fmt"
	"math"
)

// n: 配列のサイズ
// m: 和
// k: 配列
func q(n int, m int, k []int) bool {
	// 境界条件チェック
	if n < 1 || n > 50 ||
		m < 1 || m > int(math.Pow(10, 8)) {
		return false
	}
	for _, ki := range k {
		if ki < 1 || ki > int(math.Pow(10, 8)) {
			return false
		}
	}

	// kから4回取って和がmになればtrue
	for _, ki := range k {
		for _, kj := range k {
			for _, kl := range k {
				for _, km := range k {
					if ki+kj+kl+km == m {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	n := 3
	m := 10
	k := []int{1, 3, 5}
	fmt.Printf("%v", q(n, m, k))
}
