package main

import (
	"fmt"
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

func main() {
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
