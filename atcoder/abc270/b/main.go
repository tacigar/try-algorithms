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
	value, _ := strconv.Atoi(scanner.Text())
	return value
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func sign(v int) int {
	if v < 0 {
		return -1
	}
	return 1
}

func fn(x, y, z int) int {
	ax, ay, az := abs(x), abs(y), abs(z)
	sx, sy, sz := sign(x), sign(y), sign(z)
	if sx != sy || ax < ay {
		return ax
	}
	if sy == sz {
		if ay < az {
			return -1
		}
		return ax
	}
	return az*2 + ax
}

func main() {
	scanner.Split(bufio.ScanWords)
	x := scanInt()
	y := scanInt()
	z := scanInt()
	fmt.Println(fn(x, y, z))
}
