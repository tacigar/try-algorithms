package main

import "fmt"

func CountingSort(a []int, max int) []int {
	c := make([]int, max+1)
	for i := 0; i < len(a); i++ {
		c[a[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] = c[i] + c[i-1]
	}
	r := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		r[c[a[i]]-1] = a[i]
		c[a[i]]--
	}
	return r
}

func main() {
	res := []int{
		8, 1, 3, 2, 5, 6, 1, 2, 9, 9, 10, 2, 3, 5, 4, 3, 2,
	}
	fmt.Println(CountingSort(res, 10))
}
