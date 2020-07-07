package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func ScanInt() int {
	scanner.Scan()
	txt := scanner.Text()
	n, _ := strconv.Atoi(txt)
	return n
}

func main() {
	a, b, c := ScanInt(), ScanInt(), ScanInt()
	if b/a > c {
		fmt.Println(c)
	} else {
		fmt.Println(b / a)
	}
}
