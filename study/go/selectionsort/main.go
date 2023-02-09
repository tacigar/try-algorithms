package main

import "fmt"

func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := a[i]
		mi := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				mi = j
			}
		}
		a[mi], a[i] = a[i], a[mi]
	}
	return a
}

func main() {
	fmt.Println(SelectionSort([]int{
		8, 1, 3, 2, 5, 6, 1, 2, 9, 9, 10, 2, 3, 5, 4, 3, 2,
	}))
	fmt.Println(SelectionSort([]int{
		1, 2, 3, 4, 5,
	}))
	fmt.Println(SelectionSort([]int{}))
	fmt.Println(SelectionSort([]int{1}))
	fmt.Println(SelectionSort([]int{1, 2}))
	fmt.Println(SelectionSort([]int{1, 1}))
	fmt.Println(SelectionSort([]int{5, 4, 3, 2, 1}))
}
