package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n, x int
	if scanner.Scan() {
		tmp := scanner.Text()
		n, _ = strconv.Atoi(tmp)
	}
	if scanner.Scan() {
		tmp := scanner.Text()
		x, _ = strconv.Atoi(tmp)
	}

	var xs []int
	xs = append(xs, x)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			tmp := scanner.Text()
			x, _ = strconv.Atoi(tmp)
			xs = append(xs, x)
		}
	}
	sort.Ints(xs)

	var ds []int
	for i := 0; i < n; i++ {
		ds = append(ds, xs[i+1]-xs[i])
	}

	ret := ds[0]
	for i := 1; i < n; i++ {
		ret = gcd(ret, ds[i])
	}
	fmt.Println(ret)
}
