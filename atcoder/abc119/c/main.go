package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func ScanText() string {
	if scanner.Scan() {
		return scanner.Text()
	}
	fmt.Fprintf(os.Stderr, "Text not found")
	return ""
}

func ScanInt() (ret int) {
	var err error
	if scanner.Scan() {
		txt := scanner.Text()
		if ret, err = strconv.Atoi(txt); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}
	}
	return
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Min(ns ...int) int {
	if len(ns) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Arguments are necessary")
		return 0
	}
	if len(ns) == 1 {
		return ns[0]
	}
	min := ns[0]
	for i := 1; i < len(ns); i++ {
		if min > ns[i] {
			min = ns[i]
		}
	}
	return min
}

var n, a, b, c int
var ls []int

func dfs(i, aa, bb, cc int) int {
	if i == n {
		if aa > 0 && bb > 0 && cc > 0 {
			return Abs(a-aa) + Abs(b-bb) + Abs(c-cc) - 30
		}
		return math.MaxInt32
	}
	return Min(
		dfs(i+1, aa, bb, cc),
		dfs(i+1, aa+ls[i], bb, cc)+10,
		dfs(i+1, aa, bb+ls[i], cc)+10,
		dfs(i+1, aa, bb, cc+ls[i])+10,
	)
}

func main() {
	n, a, b, c = ScanInt(), ScanInt(), ScanInt(), ScanInt()
	ls = make([]int, n)
	for i := 0; i < n; i++ {
		ls[i] = ScanInt()
	}
	fmt.Println(dfs(0, 0, 0, 0))
}
