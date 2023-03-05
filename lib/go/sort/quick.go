package sort

// | 4 | 9 | 5 | 2 | 3 | 8 |(5)|
//
// (1) mi = 0, i = 0 → Exchange & mi++
// |<4>| 9 | 5 | 2 | 3 | 8 | 5 |
//
// (2) mi = 1, i = 1
// | 4 |<9>| 5 | 2 | 3 | 8 | 5 |
//
// (3) mi = 1, i = 2 → Exchange & mi++
// | 4 | 9 |<5>| 2 | 3 | 8 | 5 |
//
// (4) mi = 2, i = 3 → Exchange & mi++
// | 4 | 5 | 9 |<2>| 3 | 8 | 5 |
//
// (5) mi = 3, i = 4 → Exchange & mi++
// | 4 | 5 | 2 | 9 |<3>| 8 | 5 |
//
// (6) mi = 4, i = 5
// | 4 | 5 | 2 | 3 | 9 |<8>| 5 |
//
// (7) mi = 4, i = 6 → Exchange & mi++
// | 4 | 5 | 2 | 3 | 9 | 8 |<5>|
//
// (8) mi = 5, i = 6 → Return mi - 1
// | 4 | 5 | 2 | 3 | 5 | 8 | 9 |
func partition(data Interface, li, ri int) int {
	// TODO: Can optimize to select the pivot index.
	pi := ri

	mi := li
	for i := li; i <= ri; i++ {
		if !data.Less(pi, i) {
			data.Swap(mi, i)
			mi++
		}
	}
	return mi - 1
}

func quickSort(data Interface, li, ri int) {
	if li < ri {
		mi := partition(data, li, ri)
		quickSort(data, li, mi-1)
		quickSort(data, mi+1, ri)
	}
}

func QuickSort(data Interface) {
	quickSort(data, 0, data.Len()-1)
}
