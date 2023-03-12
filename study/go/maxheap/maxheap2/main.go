package main

import "fmt"

type Heap []int

/*
 * 1 --- 3 --- 7
 *    |     |
 *    |     |
 *    |     |
 *    |     |- 6 --- 12
 *    |
 *    |- 2 --- 5 --- 11
 *          |     |
 *          |     |- 10
 *          |
 *          |- 4 --- 9
 *                |
 *                |- 8
 */

func LeftChild(i int) int {
	return i * 2
}

func RightChild(i int) int {
	return i*2 + 1
}

func Parent(i int) int {
	return i / 2
}

func NewHeap(arr []int) Heap {
	res := make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ {
		res[i] = arr[i-1]
	}
	return res
}

func (h Heap) AdjustDown(i int) {
	l := LeftChild(i)
	r := RightChild(i)
	largestIndex := i
	if l < len(h) && h[largestIndex] < h[l] {
		largestIndex = l
	}
	if r < len(h) && h[largestIndex] < h[r] {
		largestIndex = r
	}
	if largestIndex != i {
		h[largestIndex], h[i] = h[i], h[largestIndex]
		h.AdjustDown(largestIndex)
	}
}

func (h Heap) Adjust() {
	// 二段目(上の図なら6)から下げれば十分なので÷2している。
	for i := len(h) / 2; i >= 1; i-- {
		h.AdjustDown(i)
	}
}

func (h Heap) Insert(value int) Heap {
	hh := append(h, value)

	var fn func(i int)
	fn = func(i int) {
		pi := Parent(i)
		if pi >= 1 && hh[i] > hh[pi] {
			hh[i], hh[pi] = hh[pi], hh[i]
			fn(pi)
		}
	}
	fn(len(hh) - 1)
	return hh
}

func (h Heap) DeleteRoot() Heap {
	r := h[len(h)-1]
	hh := h[:len(h)-1]
	hh[1] = r
	hh.AdjustDown(1)
	return hh
}

func main() {
	h := NewHeap([]int{4, 1, 3, 2, 9, 10, 14, 8, 7})

	h.Adjust()
	fmt.Println(h)
	hh := h.Insert(16)
	fmt.Println(hh)
	hhh := hh.DeleteRoot()
	fmt.Println(hhh)
	hhhh := hhh.DeleteRoot()
	fmt.Println(hhhh)
	h5 := hhhh.DeleteRoot()
	fmt.Println(h5)
	h6 := h5.DeleteRoot()
	fmt.Println(h6)
	h7 := h6.DeleteRoot()
	fmt.Println(h7)
}
