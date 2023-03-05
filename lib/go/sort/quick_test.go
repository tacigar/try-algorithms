package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestQuickSortIntSlice(t *testing.T) {
	input := IntSlice(make([]int, len(ints)))
	copy(input, ints)
	QuickSort(IntSlice(input))
	assert.True(t, IsSorted(input))
}
