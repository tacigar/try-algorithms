package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func ScanInt() int {
	scanner.Scan()
	txt := scanner.Text()
	n, _ := strconv.Atoi(txt)
	return n
}

type CityInfo struct{ I, P, Y, N int }

func main() {
	ScanInt()
	m := ScanInt()

	cityInfoList := make([]CityInfo, m)
	cityInfoMap := map[int][]CityInfo{}

	for i := 0; i < m; i++ {
		p := ScanInt()
		y := ScanInt()
		cityInfoList[i] = CityInfo{I: i, P: p, Y: y}
		cityInfoMap[p] = append(cityInfoMap[p], CityInfo{I: i, P: p, Y: y})
	}
	for _, cis := range cityInfoMap {
		sort.Slice(cis, func(i, j int) bool {
			return cis[i].Y < cis[j].Y
		})
		for i, ci := range cis {
			cityInfoList[ci.I].N = i + 1
		}
	}

	for i := 0; i < m; i++ {
		ci := cityInfoList[i]
		fmt.Printf("%06d%06d\n", ci.P, ci.N)
	}
}
