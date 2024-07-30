package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	var a, b [][]int = make([][]int, n+1), make([][]int, n+1)
	// a = [][]int{
	// 	{0, 0, 0, 0},
	// 	{0, 1, 2, 3},
	// 	{0, 4, 5, 6},
	// 	{0, 7, 8, 9},
	// }
	// b = [][]int{
	// 	{0, 0, 0, 0},
	// 	{0, 1, 1, 1},
	// 	{0, 1, 1, 1},
	// 	{0, 1, 1, 1},
	// }
	for i := 0; i <= n; i++ {
		a[i] = make([]int, m+1)
		b[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&b[i][j])
		}
	}
	//读数完毕
	var times int
	var x, y int
	fmt.Scan(&times)
	var bfs func(x, y int) int
	var postx []int = []int{-1, 0, 0, 1}
	var posty []int = []int{0, 1, -1, 0}
	for i := 0; i < times; i++ {

		fmt.Scanln(&x, &y)
		bfs = func(xx, yy int) int {
			count := 1
			var queue [][2]int = make([][2]int, 0)
			queue = append(queue, [2]int{xx, yy})
			for len(queue) != 0 {
				aaa, bbb := queue[0][0], queue[0][1]
				queue = queue[1:]
				for i := 0; i < 4; i++ {
					tempa, tempb := aaa+postx[i], bbb+posty[i]
					if tempa < 1 || tempa > n || tempb < 1 || tempb > m {
						continue
					}
					if a[tempa][tempb]+b[tempa][tempb] <= a[aaa][bbb] {
						queue = append(queue, [2]int{tempa, tempb})
						count = (count + 1) % 100007
					}
				}
			}
			return count
		}

		fmt.Println(bfs(x, y))
	}
}
