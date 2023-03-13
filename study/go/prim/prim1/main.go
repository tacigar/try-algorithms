package main

import (
	"fmt"
	"math"
)

func Prim(m [][]int) int {
	n := len(m)

	// cs[v]: ノードvが探索済みか否か
	cs := make([]bool, n)

	// ds[v]: すでに探索済みの頂点vと、未探索のノードの間の最小のエッジの重さ
	ds := make([]int, n)
	for i := 0; i < n; i++ {
		ds[i] = math.MaxInt
	}
	ds[0] = 0

	sum := 0

	for {
		// (ステップ1) dsをイテレートして、すでに探索済みのノードから
		// 出ている辺の内、最小の重みを持つものを見つけ出す。
		mincost := math.MaxInt
		mi := 0
		for i := 0; i < n; i++ {
			if !cs[i] && ds[i] < mincost {
				mincost = ds[i]
				mi = i
			}
		}
		// もし、新しい辺が見つからなかった場合はループを終了
		if mincost == math.MaxInt {
			break
		}
		sum += mincost
		cs[mi] = true // miは探索済みなのでマーク

		// (ステップ2) 今探索したばかりのmiから出ている辺を調べ上げて、
		// miから出ている最も重みの小さい辺を探し出し、その重みをdsに書き込む。
		// 書き込まれたdsは次のループで(ステップ1)で利用される。
		for i := 0; i < n; i++ {
			if !cs[i] && m[i][mi] != -1 {
				w := m[i][mi]
				if w < ds[i] {
					ds[i] = w
				}
			}
		}
	}
	return sum
}

func main() {
	m := [][]int{
		{-1, 2, 3, 1, -1},
		{2, -1, -1, 4, -1},
		{3, -1, -1, 1, 1},
		{1, 4, 1, -1, 3},
		{-1, -1, 1, 3, -1},
	}
	fmt.Println(Prim(m))
}
