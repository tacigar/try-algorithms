package main

import "fmt"

func Merge(a, b []int) []int {
	r := []int{}
	ai := 0
	bi := 0
	for {
		if ai >= len(a) {
			r = append(r, b[bi:]...)
			return r
		}
		if bi >= len(b) {
			r = append(r, a[ai:]...)
			return r
		}
		if a[ai] > b[bi] {
			r = append(r, b[bi])
			bi++
		} else {
			r = append(r, a[ai])
			ai++
		}
	}
}

func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	m := len(a) / 2
	a1 := MergeSort(a[:m])
	a2 := MergeSort(a[m:])
	return Merge(a1, a2)
}

func main() {
	arr := []int{3, 2, 4, 1, 5, 3, 10, 20, 4, 5, 9, 2, 6, 10, 3}
	fmt.Println(MergeSort(arr))
}
