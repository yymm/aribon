package main

import (
	"fmt"
	"math"
	"sort"
)

// n: 配列のサイズ
// m: 和
// k: 配列
func q(n int, m int, k []int) bool {
	// 境界条件チェック
	if n != len(k) {
		return false
	}
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

// 二分探索(binary search)
// m=ki+kj+kl+km を km=m-ki-kj-kl としてkmの探索を二分探索にすることで
// O(n^3logn)に高速化
func q_binary_search_for(n int, m int, k []int) bool {
	// 境界条件チェック
	if n != len(k) {
		return false
	}
	if n < 1 || n > 1000 ||
		m < 1 || m > int(math.Pow(10, 8)) {
		return false
	}
	for _, ki := range k {
		if ki < 1 || ki > int(math.Pow(10, 8)) {
			return false
		}
	}

	sort.Ints(k)

	for _, ki := range k {
		for _, kj := range k {
			for _, kl := range k {
				if binary_search_for(m-ki-kj-kl, k) == true {
					return true
				}
			}
		}
	}
	return false
}

func q_binary_search_recursive(n int, m int, k []int) bool {
	// 境界条件チェック
	if n != len(k) {
		return false
	}
	if n < 1 || n > 1000 ||
		m < 1 || m > int(math.Pow(10, 8)) {
		return false
	}
	for _, ki := range k {
		if ki < 1 || ki > int(math.Pow(10, 8)) {
			return false
		}
	}

	sort.Ints(k)

	for _, ki := range k {
		for _, kj := range k {
			for _, kl := range k {
				if binary_search_recursive(m-ki-kj-kl, k) == true {
					return true
				}
			}
		}
	}
	return false
}

func binary_search_for(v int, k []int) bool {
	mark := len(k) / 2

	for mark > 0 {
		if k[mark] == v {
			return true
		} else if k[mark] > v {
			k = k[:mark-1]
		} else {
			k = k[mark+1:]
		}
		mark = len(k) / 2
	}
	return false
}

func binary_search_recursive(v int, k []int) bool {
	mark := len(k) / 2
	if mark <= 0 {
		return false
	}
	if k[mark] == v {
		return true
	} else if k[mark] > v {
		return binary_search_recursive(v, k[:mark-1])
	} else {
		return binary_search_recursive(v, k[mark+1:])
	}
}

func test1(f func(int, int, []int) bool) {
	n := 3
	m := 10
	k := []int{1, 3, 5}
	fmt.Printf("%v\n", f(n, m, k))
}

func test2(f func(int, int, []int) bool) {
	n := 3
	m := 9
	k := []int{1, 3, 5}
	fmt.Printf("%v\n", f(n, m, k))
}

func main() {
	// 4重ループ
	test1(q)
	test2(q)
	// O(n^3logn)
	// binary_search_for
	test1(q_binary_search_for)
	test2(q_binary_search_for)
	// O(n^3logn)
	// binary_search_recursive
	test1(q_binary_search_recursive)
	test2(q_binary_search_recursive)
}
