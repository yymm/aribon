package main

import (
	"fmt"
)

func q1(l int, x []int, d []bool) int {
	s := 0
	finish := 0
	for {
		// アリをひっくり返す
		for i := 0; i < len(x); i++ {
			p := 0
			if d[i] == true {
				p = x[i] + 1
			} else {
				p = x[i] - 1
			}
			for j := i + 1; j < len(x); j++ {
				if p == x[j] && d[j] != d[i] {
					d[i] = !d[i]
					d[j] = !d[j]
				}
			}
		}
		// アリをすすめる
		for i := 0; i < len(x); i++ {
			if x[i] < 0 || x[i] > l {
				continue
			}
			if d[i] == true {
				x[i] += 1
			} else {
				x[i] -= 1
			}
			if x[i] < 0 || x[i] > l {
				finish++
			}
		}
		if finish == len(x) {
			break
		}
		s++
	}
	return s
}

func r(p int, d []bool, l int, x []int) {
	if p == 0 {
		xc := make([]int, len(x))
		copy(xc, x)
		dc := make([]bool, len(d))
		copy(dc, d)
		fmt.Println(q1(l, xc, dc), dc)
		return
	}
	d[len(d)-p] = false
	r(p-1, d, l, x)
	d[len(d)-p] = true
	r(p-1, d, l, x)
}

func q(l int, x []int) (int, []int) {
	d := make([]bool, len(x)) // "false" fill

	r(len(d), d, l, x)

	return l, x
}

func main() {
	l := 10
	x := []int{2, 6, 7}
	q(l, x)
}
