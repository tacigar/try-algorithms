package main

import "fmt"

/*
 * pivot = 4
 *
 * (0) i = 0, mi = 0
 * |<4>| 2 | 5 | 9 | 1 | 7 | 8 | 4 |
 *
 * 4 <= 4なのでiとmiを交換し、mi++
 *
 * (1) i = 1, mi = 1
 * | 4 |<2>| 5 | 9 | 1 | 7 | 8 | 4 |
 *
 * 2 <= 4なのでiとmiを交換し、mi++
 *
 * (2) i = 2, mi = 2
 * | 4 | 2 |<5>| 9 | 1 | 7 | 8 | 4 |
 *
 * (3) i = 3, mi = 2
 * | 4 | 2 | 5 |<9>| 1 | 7 | 8 | 4 |
 *
 * (4) i = 4, mi = 2
 * | 4 | 2 | 5 | 9 |<1>| 7 | 8 | 4 |
 *
 * 1 <= 4なのでiとmiを交換し、mi++
 *
 * (5) i = 5, mi = 3
 * | 4 | 2 | 1 | 9 | 5 |<7>| 8 | 4 |
 *
 * (6) i = 6, mi = 3
 * | 4 | 2 | 1 | 9 | 5 | 7 |<8>| 4 |
 *
 * (7) i = 7, mi = 3
 * | 4 | 2 | 1 | 9 | 5 | 7 | 8 |<4>|
 *
 * 4 <= 4なのでiとmiを交換し、mi++
 *
 * (8) i = 8, mi = 4
 * | 4 | 2 | 1 | 4 | 5 | 7 | 8 | 9 |
 *
 * このとき真ん中の4は必ず正しい位置に配置されていることに注意。
 * 要するに0-3と5-7を調べていけば良い。
 */
func Partition(a []int, li, ri int) int {
	p := a[ri]
	mi := li
	for i := li; i <= ri; i++ {
		if a[i] <= p {
			a[i], a[mi] = a[mi], a[i]
			mi++
		}
	}
	return mi - 1
}

func QuickSort(a []int, li, ri int) {
	if ri <= li {
		return
	}
	mi := Partition(a, li, ri)
	QuickSort(a, li, mi-1)
	QuickSort(a, mi+1, ri)
}

func main() {
	arr := []int{3, 2, 4, 1, 5, 3, 10, 20, 4, 5, 9, 2, 6, 10, 3}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
