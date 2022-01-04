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
	if n >= 42 {
		n++
	}
	n1 := n % 10
	n2 := (n / 10) % 10
	fmt.Println("AGC0" + strconv.Itoa(n2) + strconv.Itoa(n1))
}
