package sort

func BubbleSort(data Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
