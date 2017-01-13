package main

import (
	"fmt"
)

func q(n int, a []int) int {
	if n < 3 {
		return 0
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if a[i]+a[j] < a[k] || a[j]+a[k] < a[i] || a[k]+a[i] < a[j] {
					continue
				}
				temp := a[i] + a[j] + a[k]
				if max < temp {
					max = temp
				}
			}
		}
	}
	return max
}

func main() {
	n := 5
	a := []int{2, 3, 4, 5, 10}
	fmt.Printf("%d\n", q(n, a))
}
