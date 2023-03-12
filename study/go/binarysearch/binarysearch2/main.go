package main

import "fmt"

func BinarySearch(a []int, t int) bool {
	l := 0
	r := len(a) - 1
	for {
		if r < l {
			return false
		}
		i := (l + r) / 2
		if a[i] == t {
			return true
		}
		if a[i] > t {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

func main() {
	fmt.Println(BinarySearch([]int{
		1, 2, 3, 4, 5,
	}, 3))
	fmt.Println(BinarySearch([]int{
		1, 2, 3, 4, 5,
	}, 4))
	fmt.Println(BinarySearch([]int{
		1, 2, 3, 4, 5,
	}, 1))
	fmt.Println(BinarySearch([]int{
		1, 2, 3, 4, 5,
	}, 6))
	fmt.Println(BinarySearch([]int{
		1, 2, 3, 4, 5,
	}, 1))
	fmt.Println(BinarySearch([]int{
		1, 2, 3, 4, 5,
	}, 5))
}
