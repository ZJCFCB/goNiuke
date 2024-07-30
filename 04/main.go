package main

import (
	"fmt"
)

func main() {
	var n, m, t int
	fmt.Scanln(&n, &m, &t)
	var graph map[int][]int = make(map[int][]int)
	var from, to int
	var indegree []int = make([]int, n+1)
	var queue []int = make([]int, n+1)
	var tp []int = make([]int, n+1)
	//读图+计算节点的入度

	for i := 0; i < m; i++ {
		fmt.Scanln(&from, &to)
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		tp = append(tp, point)
		for _, vv := range graph[point] {
			indegree[vv]--
			if indegree[vv] == 0 {
				queue = append(queue, vv)
			}
		}
	}
	var dp []int = make([]int, n+1)
	dp[t] = 1

	for i := len(dp) - 1; i >= 0; i-- {
		for _, vv := range graph[i] {
			dp[i] = (dp[i] + dp[vv]) % 100007
		}
	}
	var times int
	var start int
	fmt.Scanln(&times)

	for i := 0; i < times; i++ {
		fmt.Scanln(&start)
		if start < 1 || start > n {
			fmt.Println(0)
			continue
		}
		fmt.Println(dp[start])
	}
}
