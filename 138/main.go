package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	var n, m, j, k int
	var MAX int = 2010
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	fmt.Scanln(&j, &k)
	var graph, weight [][]int = make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		weight[i] = make([]int, m)
		for j := 0; j < m; j++ {
			weight[i][j] = MAX
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
	var postx []int = []int{0, 0, 1, -1}
	var posty []int = []int{-1, 1, 0, 0}

	weight[j][k] = 0
	var queue [][]int
	queue = append(queue, []int{j, k})
	for len(queue) != 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			tempx, tempy := x+postx[i], y+posty[i]

			if tempx >= 0 && tempx < n && tempy >= 0 && tempy < m && graph[tempx][tempy] != 0 {
				if weight[tempx][tempy] > graph[x][y]+weight[x][y] {
					weight[tempx][tempy] = graph[x][y] + weight[x][y]
					queue = append(queue, []int{tempx, tempy})
				}
				// weight[tempx][tempy] = min(weight[tempx][tempy], graph[x][y]+weight[x][y])
				// queue = append(queue, []int{tempx, tempy})

			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] != 0 {
				ans = max(ans, weight[i][j])
			}
		}
	}
	if ans == MAX {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
