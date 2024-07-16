package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	var wi []int = make([]int, n)
	var ai []int = make([]int, n)
	var ans int
	for i := 0; i < n; i++ {
		fmt.Scan(&wi[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&ai[i])
	}
	var dp []int = make([]int, n)
	copy(dp, wi)
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			post := i + ai[i]*j
			if post >= n {
				break
			}
			dp[post] = max(dp[post], dp[i]+wi[post])
			ans = max(ans, dp[post])
		}
		fmt.Println(dp)
	}
	fmt.Println(ans)
}
