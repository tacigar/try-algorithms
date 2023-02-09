package main

import "fmt"

func ShellSort(a []int) []int {
	n := len(a)
	hs := []int{8, 4, 2, 1}
	for h := range hs {
		for d := 0; d < h; d += 1 {
			for i := d; i < n; i += h {
				e := a[i]
				for j := i - h; j >= 0; j -= h {
					if a[j] < e {
						a[j+h] = e
						break
					}
					a[j+h] = a[j]
					a[j] = e
				}
			}
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
