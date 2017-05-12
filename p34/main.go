package main

import (
	"fmt"
)

// 深さ優先探索DFS
func dfs(i int, sum int, k int, a []int) bool {
	if i == len(a) {
		return sum == k
	}
	if dfs(i+1, sum, k, a) == true {
		return true
	}
	if dfs(i+1, sum+a[i], k, a) == true {
		return true
	}
	return false
}

func q(k int, a []int) bool {
	return dfs(0, 0, k, a)
}

func main() {
	a := []int{3, 2, 1, 5, 4}
	fmt.Println(q(17, a))
}
