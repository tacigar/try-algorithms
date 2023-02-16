package main

import "fmt"

func Partition(a []int, l, r int) int {
	p := a[r]
	m := l
	for i := l; i <= r; i++ {
		if a[i] <= p {
			a[i], a[m] = a[m], a[i]
			m++
		}
	}
	return m - 1
}

func QuickSortImpl(a []int, l, r int) {
	if l >= r {
		return
	}
	m := Partition(a, l, r)
	QuickSortImpl(a, l, m-1)
	QuickSortImpl(a, m+1, r)
}

func QuickSort(a []int) {
	QuickSortImpl(a, 0, len(a)-1)
}

func main() {
	arr := []int{3, 2, 4, 1, 5, 3, 10, 20, 4, 5, 9, 2, 6, 10, 3}
	QuickSort(arr)
	fmt.Println(arr)
}
