package main

import "fmt"

func CountingSort(a []int, k int) []int {
	c := make([]int, k+1)
	b := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		c[a[i]]++
	}
	for i := 1; i <= k; i++ {
		c[i] += c[i-1]
	}
	for i := len(a) - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]]--
	}
	return b
}

func main() {
	fmt.Println(CountingSort([]int{8, 1, 3, 2, 5, 6, 1, 2, 9, 9, 10, 2, 3, 5, 4, 3, 2}, 11))
}
