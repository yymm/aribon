package main

import (
	"fmt"
	"time"
)

func dfs(n, m, i, j int, garden [][]int) {
	garden[i][j] = 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if i+di >= 0 && i+di < n && j+dj >= 0 && j+dj < m && garden[i+di][j+dj] == 1 {
				dfs(n, m, i+di, j+dj, garden)
			}
		}
	}
}

func q(n, m int, garden [][]int) int {
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if garden[i][j] == 1 {
				dfs(n, m, i, j, garden)
				count++
			}
		}
	}
	return count
}

// うねうねした池を作る
// 5*5の例
// 11111
// 00001
// 11111
// 10000
// 11111
func uneune(n int) [][]int {
	garden := make([][]int, n)
	for i := 0; i < n; i++ {
		garden[i] = make([]int, n)
		for j := 0; j < n; j++ {
			v := 1
			if i%2 == 1 {
				if j == n-1 && (i%3 == 1 || i%3 == 2) {
					v = 1
				} else if j == 0 && i%3 == 0 {
					v = 1
				} else {
					v = 0
				}
			}
			garden[i][j] = v
			//fmt.Printf("%d", v)
		}
		//fmt.Printf("\n")
	}
	return garden
}

// 再帰の関数呼び出しがstackoverflowになるか確認するテスト
// goだと1000でもいけた、i5-4200U(1.4GHz*4)で0.2秒くらい
func test_uneune() {
	n := 1000
	start := time.Now()

	fmt.Println(q(n, n, uneune(n)))

	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
}

// 🐜本のテスト
func test() {
	n := 10
	m := 12
	garden := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	}
	fmt.Println(q(n, m, garden))
}

func main() {
	test_uneune()
}
