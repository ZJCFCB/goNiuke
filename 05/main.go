package main

import (
	"fmt"
	"sort"
)

type point struct {
	x, y int
}
type edge struct {
	p1, p2 int
	weight int
}

func distval(xx, yy point) int {
	return (xx.x-yy.x)*(xx.x-yy.x) + (xx.y-yy.y)*(xx.y-yy.y)
}

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	var p []point = make([]point, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&p[i].x, &p[i].y)
	}
	var e []edge = make([]edge, 0)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			e = append(e, edge{p1: i, p2: j, weight: distval(p[i], p[j])})
		}
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].weight < e[j].weight
	})

	var father []int = make([]int, n+1)
	var nums []int = make([]int, n+1)

	//注意这里的nums，以及后面的初始化

	for i := 0; i <= n; i++ {
		father[i] = i
		nums[i] = 1
	}

	var find func(x int) int
	find = func(x int) int {
		if x != father[x] {
			father[x] = find(father[x])
		}
		return father[x]
	}

	for _, edg := range e {
		v1 := find(edg.p1)
		v2 := find(edg.p2)
		if v1 != v2 {
			father[v1] = v2
			nums[v2] += nums[v1]
			if nums[v2] >= k {
				fmt.Println(edg.weight)
				return
			}
		}
	}
}
