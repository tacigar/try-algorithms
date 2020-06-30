package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, 0)
	scanner.Buffer(buf, 10000000000)
	scanner.Split(bufio.ScanWords)
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

func main() {
	n := ScanInt()
	a, b, c, d, e := ScanInt(), ScanInt(), ScanInt(), ScanInt(), ScanInt()
	m := Min(a, b, c, d, e)
	if n%m == 0 {
		fmt.Println(n/m + 4)
	} else {
		fmt.Println(n/m + 5)
	}
}
