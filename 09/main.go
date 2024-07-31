package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var data []int = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&data[i])
	}
	sort.Ints(data)
	var ans int
	for i := n; i > 0; i-- {
		ans += i * data[i]
	}
	fmt.Println(ans)
}
