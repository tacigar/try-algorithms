package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int64 {
	scanner.Scan()
	value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return value
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	for x := 0; x < 998244353; x++ {
		if (n-int64(x))%int64(998244353) == 0 {
			fmt.Println(x)
			return
		}
	}
}
