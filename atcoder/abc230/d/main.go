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
	a, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return a
}

type wall struct {
	l      int
	r      int
	broken bool
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := scanInt()
	d := scanInt()
	m := map[int]map[int]struct{}{}
	ws := make([]wall, n)
	for i := 0; i < n; i++ {
		ws[i].l = scanInt()
		ws[i].r = scanInt()
		ws[i].broken = false
		for j := ws[i].l; j <= ws[i].r; j++ {
			if _, ok := m[j]; !ok {
				m[j] = map[int]struct{}{}
			}
			m[j][i] = struct{}{}
		}
	}
	sort.Slice(ws, func(i, j int) bool {
		return ws[i].r < ws[j].r
	})
	cnt := 0
	for i := 0; i < n; i++ {
		if ws[i].broken {
			continue
		}
		cnt++
		for j := ws[i].r; j < ws[i].r+d; j++ {
			for k := range m[j] {
				ws[k].broken = true
			}
		}
	}
	fmt.Println(cnt)
}
