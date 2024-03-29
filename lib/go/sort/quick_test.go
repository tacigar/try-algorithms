package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortIntSlice(t *testing.T) {
	input := IntSlice(make([]int, len(ints)))
	copy(input, ints)
	QuickSort(IntSlice(input))
	assert.True(t, IsSorted(input))
}
