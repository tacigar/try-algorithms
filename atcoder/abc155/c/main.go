package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanText() string {
	scanner.Scan()
	return scanner.Text()
}

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
	d := map[string]int{}
	for i := 0; i < n; i++ {
		s := scanText()
		d[s]++
	}

	ret := []string{}
	max := 0
	for k, v := range d {
		if v > max {
			max = v
			ret = []string{k}
		} else if v == max {
			ret = append(ret, k)
		}
	}
	sort.Strings(ret)

	for _, k := range ret {
		fmt.Println(k)
	}
}
