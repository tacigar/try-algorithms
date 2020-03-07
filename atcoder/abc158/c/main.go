package main

import (
	"bufio"
	"fmt"
	"math"
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

	a := scanInt()
	b := scanInt()

	for x := 0; x < 1015; x++ {
		t1 := int(math.Floor(float64(x) * 0.08))
		t2 := int(math.Floor(float64(x) * 0.1))
		if t1 == a && t2 == b {
			fmt.Println(x)
			return
		}
	}
	fmt.Println(-1)
}
