package main

import (
	"fmt"
)

func main() {
	// 入力
	var N int
	fmt.Scan(&N)

	// 奇数なら枝刈り
	if N%2 != 0 {
		return
	}

	/**
	例: N=3
	000
	001
	010
	011
	100
	101
	110
	111
	*/
	for s:=0;s<(1<<N);s++ { // 2^N-1通りの文字列をすべて生成
		match := 0
		candidate := ""
		// 一文字ずつ見ていく。0始まりだと右端bitから見ることになってしまい辞書順が崩れるので、左端bitから見るためにN-1から見る
		for b :=N-1; b >=0; b-- {
			if (s & (1 << b)) == 0 { // sのうちi番目が0かどうか 0なら( 1なら) とする
				match++
				candidate += "("
			} else {
				match--
				candidate += ")"
			}
			if match < 0 { // 途中で負になったら閉じカッコが多すぎる正しくないカッコ列
				break
			}
		}
		if match == 0 {
			fmt.Println(candidate)
		}
	}
}
