package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int {
	scanner.Scan()
	a, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt()
	k := scanInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			ret[i] = float64(a[i]+1) * 0.5
		} else if i >= k {
			ret[i] = ret[i-1] - float64(a[i-k]+1)*0.5 + float64(a[i]+1)*0.5
		} else {
			ret[i] = ret[i-1] + float64(a[i]+1)*0.5
		}
	}

	max := float64(0.0)
	for _, v := range ret {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)
}
