package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)
	fmt.Scan(&q)
	var data []int = make([]int, n+2)
	var color []byte = []byte{'R', 'G', 'B'}
	var l, r int
	for i := 0; i < q; i++ {
		fmt.Scanln(&l, &r)
		data[l]++
		data[r+1]--
	}
	for i := 1; i <= n; i++ {
		data[i] += data[i-1]
		if data[i] == 0 {
			fmt.Print("O")
		} else {
			fmt.Print("%c", color[data[i]%3])
		}
	}
}
