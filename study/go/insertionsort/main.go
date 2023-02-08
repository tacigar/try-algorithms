package main

import "fmt"

func InsertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		e := a[i]
		for j := i - 1; j >= 0; j-- {
			if e >= a[j] {
				a[j+1] = e
				break
			}
			a[j+1] = a[j]
			a[j] = e
		}
	}
	return a
}

func main() {
	fmt.Println(InsertionSort([]int{
		8, 1, 3, 2, 5, 6, 1, 2, 9, 9, 10, 2, 3, 5, 4, 3, 2,
	}))
	fmt.Println(InsertionSort([]int{
		1, 2, 3, 4, 5,
	}))
	fmt.Println(InsertionSort([]int{}))
	fmt.Println(InsertionSort([]int{1}))
	fmt.Println(InsertionSort([]int{1, 2}))
	fmt.Println(InsertionSort([]int{1, 1}))
	fmt.Println(InsertionSort([]int{5, 4, 3, 2, 1}))
}
