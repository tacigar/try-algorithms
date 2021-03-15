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

type employee struct {
	index int
	data  []int
}

func changeOrder(x1, x2, xi employee, index int) (employee, employee) {
	if xi.data[index] <= x1.data[index] {
		return xi, x1
	}
	if xi.data[index] < x2.data[index] {
		return x1, xi
	}
	return x1, x2
}

func newEmployee() employee {
	return employee{
		index: -1,
		data:  []int{math.MaxInt32, math.MaxInt32},
	}
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	a1, a2, b1, b2 := newEmployee(), newEmployee(), newEmployee(), newEmployee()
	for i := 0; i < n; i++ {
		e := employee{i, []int{scanInt(), scanInt()}}
		a1, a2 = changeOrder(a1, a2, e, 0)
		b1, b2 = changeOrder(b1, b2, e, 1)
	}
	if a1.index != b1.index {
		fmt.Println(int(math.Max(float64(a1.data[0]), float64(b1.data[1]))))
	} else {
		c1 := a1.data[0] + a1.data[1]
		c2 := math.Max(float64(a1.data[0]), float64(b2.data[1]))
		c3 := math.Max(float64(a2.data[0]), float64(b1.data[1]))
		m := math.Min(math.Min(float64(c1), c2), c3)
		fmt.Println(m)
	}
}
