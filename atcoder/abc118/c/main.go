package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GCD(a, b int) int {
	if a%b == 0 {
		return b
	}
	return GCD(b, a%b)
}

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
	n := ScanInt()
	gcd := ScanInt()
	for i := 1; i < n; i++ {
		gcd = GCD(gcd, ScanInt())
	}
	fmt.Println(gcd)
}
