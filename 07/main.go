package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func main() {
	var n, q int
	fmt.Scanln(&n, &q)
	var draw []int = make([]int, n+1)
	var change []int = make([]int, n+1)
	// for i := 0; i <= n; i++ {
	// 	draw[i] = 0
	// }
	for i := 0; i < q; i++ {
		var key int
		fmt.Scan(&key)
		switch key {
		case 1:
			var l, r, num int
			fmt.Scan(&l, &r, &num)
			var color []int = make([]int, num)
			for j := 0; j < num; j++ {
				fmt.Scan(&color[j])
			}
			for z := 0; l+z <= r; z++ {
				newcolor := color[z%num]
				if newcolor != draw[l+z] {
					change[l+z]++
				}
				draw[l+z] = newcolor
			}
		case 2:
			var post int
			fmt.Scan(&post)
			fmt.Println(change[post], draw[post])
		case 3:
			var l, r int
			fmt.Scan(&l, &r)
			var maxnum int
			var count map[int]int = make(map[int]int)
			for i := l; i <= r; i++ {
				count[draw[i]]++
				maxnum = max(maxnum, count[draw[i]])
			}
			fmt.Println(maxnum)
		}
	}
}
