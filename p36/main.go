package main

import (
	"fmt"
	"sort"
)

func search(x, y, step int, maze [][]int, reaching *[]int) {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || nx > len(maze)-1 ||
			ny < 0 || ny > len(maze[x])-1 ||
			maze[nx][ny] == 1 {
			continue
		}
		if maze[nx][ny] == 3 {
			*reaching = append(*reaching, step)
			return
		}
		if maze[nx][ny] == 0 {
			step++
			maze[nx][ny] = 1
			search(nx, ny, step, maze, reaching)
		}
	}
}

func q(n, m int, maze [][]int) int {
	reaching := []int{}
	// スタート位置に移動
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if maze[i][j] == 2 {
				search(i, j, 0, maze, &reaching)
				break
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(maze[i][j])
		}
		fmt.Printf("\n")
	}
	sort.Ints(reaching)
	fmt.Println(reaching)
	return reaching[0]
}

func test() {
	n := 10
	m := 10
	// 0:道 1:壁 2:スタート 3:ゴール
	maze := [][]int{
		{1, 2, 1, 1, 1, 1, 1, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		{0, 1, 0, 1, 1, 0, 1, 1, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 3, 1},
	}
	fmt.Print(q(n, m, maze))
}

func test_simple() {
	n := 5
	m := 5
	// 0:道 1:壁 2:スタート 3:ゴール
	maze := [][]int{
		{1, 2, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{1, 1, 0, 3, 1},
	}
	fmt.Print(q(n, m, maze))
}

func main() {
	//test_simple()
	test()
}
