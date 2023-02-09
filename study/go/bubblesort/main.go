package main

import "fmt"

func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		flag := false
		for j := len(a) - 2; j >= i; j-- {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return a
}

func main() {
	fmt.Println(BubbleSort([]int{
		8, 1, 3, 2, 5, 6, 1, 2, 9, 9, 10, 2, 3, 5, 4, 3, 2,
	}))
	fmt.Println(BubbleSort([]int{
		1, 2, 3, 4, 5,
	}))
	fmt.Println(BubbleSort([]int{}))
	fmt.Println(BubbleSort([]int{1}))
	fmt.Println(BubbleSort([]int{1, 2}))
	fmt.Println(BubbleSort([]int{1, 1}))
	fmt.Println(BubbleSort([]int{5, 4, 3, 2, 1}))

}
