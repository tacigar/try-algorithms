package main

import "fmt"

func merge(arr []int, l, m, r int) {
	bsize := r - l + 1
	buffer := make([]int, bsize)
	bi := 0

	li := l
	mi := m

	for {
		if li >= m {
			for ; bi < bsize; bi++ {
				buffer[bi] = arr[mi]
				mi++
			}
			break
		}

		if mi > r {
			for ; bi < bsize; bi++ {
				buffer[bi] = arr[li]
				li++
			}
			break
		}

		if arr[li] < arr[mi] {
			buffer[bi] = arr[li]
			li++
		} else {
			buffer[bi] = arr[mi]
			mi++
		}
		bi++
	}

	for i := 0; i < bsize; i++ {
		arr[i+l] = buffer[i]
	}
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m+1, r)
	}
}

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	MergeSort(arr)
	fmt.Println(arr)
}
