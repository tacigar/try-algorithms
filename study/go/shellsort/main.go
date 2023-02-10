package main

import "fmt"

func InsertionSort(a []int, d, s int) []int {
	n := len(a)
	for i := s; i < n; i += d {
		e := a[i]
		for j := i - d; j >= 0; j -= d {
			if a[j] < e {
				a[j+d] = e
				break
			}
			a[j+d], a[j] = a[j], a[j+d]
		}
	}
	return a
}

func ShellSort(a []int) []int {
	ds := []int{8, 4, 2, 1}
	for d := range ds {
		for s := 0; s < d; s++ {
			InsertionSort(a, d, s)
		}
	}
	return a
}

func main() {
	fmt.Println(ShellSort([]int{
		8, 1, 3, 2, 5, 6, 1, 2, 9, 9, 10, 2, 3, 5, 4, 3, 2,
	}))
	fmt.Println(ShellSort([]int{
		1, 2, 3, 4, 5,
	}))
	fmt.Println(ShellSort([]int{}))
	fmt.Println(ShellSort([]int{1}))
	fmt.Println(ShellSort([]int{1, 2}))
	fmt.Println(ShellSort([]int{1, 1}))
	fmt.Println(ShellSort([]int{5, 4, 3, 2, 1}))
}
