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

type To struct {
	e int
	w int
}

func makeLengths(n int, to map[int][]To, l []int) {
	l[0] = 0
	var impl func(p, gp int)
	impl = func(p, gp int) {
		if to[p] == nil {
			return
		}
		for _, nextTo := range to[p] {
			if nextTo.e == gp {
				continue
			}
			l[nextTo.e] = nextTo.w + l[p]
			impl(nextTo.e, p)
		}
	}
	impl(0, -1)
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt()

	to := map[int][]To{}

	for i := 0; i < n-1; i++ {
		u := scanInt()
		v := scanInt()
		w := scanInt()

		if to[u-1] == nil {
			to[u-1] = []To{{e: v - 1, w: w}}
		} else {
			to[u-1] = append(to[u-1], To{e: v - 1, w: w})
		}

		if to[v-1] == nil {
			to[v-1] = []To{{e: u - 1, w: w}}
		} else {
			to[v-1] = append(to[v-1], To{e: u - 1, w: w})
		}
	}

	l := make([]int, n)
	makeLengths(n, to, l)

	for i := range l {
		if l[i]%2 == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
