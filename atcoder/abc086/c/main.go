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
	if v > 0 {
		return v
	}
	return -v
}

type Datum struct {
	t int
	x int
	y int
}

func can(prev, next Datum) bool {
	dt := next.t - prev.t
	dd := abs(next.x-prev.x) + abs(next.y-prev.y)
	if dd > dt {
		return false
	}
	return (dt-dd)%2 == 0
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	data := make([]Datum, n+1)
	data[0] = Datum{0, 0, 0}
	for i := 1; i < n+1; i++ {
		data[i].t = scanInt()
		data[i].x = scanInt()
		data[i].y = scanInt()
	}
	for i := 0; i < n; i++ {
		if !can(data[i], data[i+1]) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
