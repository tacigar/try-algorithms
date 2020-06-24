package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func ScanInt() int {
	scanner.Scan()
	t := scanner.Text()
	n, _ := strconv.Atoi(t)
	return n
}

func main() {
	n, m := ScanInt(), ScanInt()
	if n >= m {
		fmt.Println(0)
		return
	}
	xs := make([]int, m)
	for i := 0; i < m; i++ {
		xs[i] = ScanInt()
	}
	sort.Ints(xs)
	ds := make([]int, m-1)
	for i := 0; i < m-1; i++ {
		ds[i] = xs[i+1] - xs[i]
	}
	ret := 0
	sort.Ints(ds)
	for i := 0; i < m-1-(n-1); i++ {
		ret += ds[i]
	}
	fmt.Println(ret)
}
