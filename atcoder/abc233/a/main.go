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
	x := scanInt()
	y := scanInt()
	if y < x {
		fmt.Println(0)
	} else if (y-x)%10 != 0 {
		fmt.Println((y-x)/10 + 1)
	} else {
		fmt.Println((y - x) / 10)
	}
}
