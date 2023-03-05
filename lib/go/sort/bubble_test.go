package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortIntSlice(t *testing.T) {
	input := IntSlice(make([]int, len(ints)))
	copy(input, ints)
	BubbleSort(IntSlice(input))
	assert.True(t, IsSorted(input))
}
