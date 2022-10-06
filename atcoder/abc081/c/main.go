package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int {
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	return value
}

type P struct {
	v int
	c int
}

type Ps []P

func (ps Ps) Len() int {
	return len(ps)
}

func (ps Ps) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps Ps) Less(i, j int) bool {
	return ps[i].c < ps[j].c
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	am := map[int]int{}
	for i := 0; i < n; i++ {
		v := scanInt()
		am[v] = am[v] + 1
	}

	var ps Ps
	for k, v := range am {
		ps = append(ps, P{v: k, c: v})
	}
	sort.Sort(ps)

	lam := len(am)

	if lam <= k {
		fmt.Println(0)
	} else if lam > k {
		cnt := 0
		for i := 0; i < lam-k; i++ {
			cnt += ps[i].c
		}
		fmt.Println(cnt)
	}
}
