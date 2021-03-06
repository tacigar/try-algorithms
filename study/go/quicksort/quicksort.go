package main

import (
	"fmt"
	"math/rand"
)

func quicksort(arr []int) {
	var quicksortImpl func(l, r int)
	quicksortImpl = func(l, r int) {
		if l >= r {
			return
		}
		pi := rand.Intn(r-l+1) + l
		pivot := arr[pi]

		li, ri := l, r
		for {
			for ; arr[li] < pivot && li < r; li++ {
			}
			for ; arr[ri] > pivot && ri > l; ri-- {
			}
			if li < ri {
				arr[li], arr[ri] = arr[ri], arr[li]
				ri--
				li++
			} else {
				quicksortImpl(l, li-1)
				quicksortImpl(li, r)
				break
			}
		}
	}
	quicksortImpl(0, len(arr)-1)
}

func main() {
	arr := []int{3, 2, 4, 1, 5, 3, 10, 20, 4, 5, 9, 2, 6, 10, 3}
	quicksort(arr)
	fmt.Println(arr)
}
