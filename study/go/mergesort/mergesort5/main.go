package main

import "fmt"

func merge(arr []int, l, m, r int) {
	bsize := r - l + 1
	buf := make([]int, bsize)
	bi, li, mi := 0, l, m

	for {
		// 左側のイテレータが中間を超えている場合、右側を全部追加してbreak
		if li >= m {
			for ; bi < bsize; bi++ {
				buf[bi] = arr[mi]
				mi++
			}
			break
		}
		// 右側のイテレータが右端を超えている場合、左側を全部追加してbreak
		if mi > r {
			for ; bi < bsize; bi++ {
				buf[bi] = arr[li]
				li++
			}
			break
		}
		// liとmiの小さい方を追加して、biをインクリメント
		if arr[li] < arr[mi] {
			buf[bi] = arr[li]
			li++
		} else {
			buf[bi] = arr[mi]
			mi++
		}
		bi++
	}
	for i := 0; i < bsize; i++ {
		arr[l+i] = buf[i]
	}
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		// ここでm+1を渡すことが非常に重要
		// もしmを渡してしまった場合、右側のイテレータと左側のイテレータが重なるので挙動がおかしくなる
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
