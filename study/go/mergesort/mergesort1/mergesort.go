package main

import "fmt"

func merge(arr1, arr2 []int) []int {
	ret := []int{}
	i, j := 0, 0
	for {
		if i == len(arr1) {
			ret = append(ret, arr2[j:]...)
			break
		}
		if j == len(arr2) {
			ret = append(ret, arr1[i:]...)
			break
		}
		if arr1[i] < arr2[j] {
			ret = append(ret, arr1[i])
			i++
		} else {
			ret = append(ret, arr2[j])
			j++
		}
	}
	return ret
}

func mergesort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	m := len(arr) / 2
	a1 := mergesort(arr[0:m])
	a2 := mergesort(arr[m:])
	return merge(a1, a2)
}

func main() {
	arr := []int{3, 2, 4, 1, 5, 3, 10, 20, 4, 5, 9, 2, 6, 10, 3}
	fmt.Println(mergesort(arr))
}
