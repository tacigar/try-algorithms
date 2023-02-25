package main

import "fmt"

func Parent(i int) int {
	return i / 2
}

func Left(i int) int {
	return 2 * i
}

func Right(i int) int {
	return 2*i + 1
}

func MaxHeapify(h []int, i, size int) {
	l := Left(i)
	r := Right(i)
	largest := i
	if l <= size && h[l-1] > h[i-1] {
		largest = l
	}
	if r <= size && h[r-1] > h[largest-1] {
		largest = r
	}
	if largest != i {
		h[i-1], h[largest-1] = h[largest-1], h[i-1]
		MaxHeapify(h, largest, size)
	}
}

func BuildMaxHeap(h []int, size int) {
	for i := size / 2; i >= 1; i-- {
		MaxHeapify(h, i, size)
	}
}

func main() {
	h := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	BuildMaxHeap(h, len(h))
	fmt.Println(h)
}
