package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n, k int
	if scanner.Scan() {
		tmp := scanner.Text()
		n, _ = strconv.Atoi(tmp)
	}
	if scanner.Scan() {
		tmp := scanner.Text()
		k, _ = strconv.Atoi(tmp)
	}

	var xs []int
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			tmp := scanner.Text()
			x, _ := strconv.Atoi(tmp)
			xs = append(xs, x)
		}
	}

	min := math.MaxInt32
	for i := 0; i < n-k+1; i++ {
		var v int
		if xs[i] >= 0 && xs[i+k-1] >= 0 {
			v = xs[i+k-1]
		} else if xs[i] < 0 && xs[i+k-1] > 0 {
			if -xs[i] < xs[i+k-1] {
				v = xs[i+k-1] - 2*xs[i]
			} else {
				v = 2*xs[i+k-1] - xs[i]
			}
		} else {
			v = -xs[i]
		}
		if min > v {

			min = v
		}
	}
	fmt.Println(min)
}
