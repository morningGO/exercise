package main

import (
	"math/rand"
	"sort"
	"time"

	"github.com/kr/pretty"
)

func main() {
	/*
	   * 配列の連結
	     * 任意の配列を複数作成し、その配列を連結し画面に表示する
	*/
	abcdefghi := [...]string{"abc", "def", "ghi"}
	lmnopq := [...]string{"lmn", "opq"}

	pretty.Printf("#1: %# v\n", append(abcdefghi[:], lmnopq[:]...))

	/*
	   * 配列のスライス
	     * 任意の配列を作成し、配列をスライスし画面に表示する
	     * スライスについて：http://jxck.hatenablog.com/entry/golang-slice-internals
	     * ってか a tour of goでやったよね？
	*/
	pretty.Printf("#2 :%# v , %# v\n", abcdefghi[0:1], abcdefghi[1:3])

	/*
	   * 配列値のソート
	     * 任意の配列を作成し、値に整数値を代入します。代入した値をソートしてまた配列に代入し画面に表示する
	*/
	numbers := [5]int{}
	rand.Seed(time.Now().Unix())
	for i := range numbers {
		numbers[i] = rand.Intn(10)
	}
	sort.Sort(sort.IntSlice(numbers[:]))
	pretty.Printf("#3 :%# v\n", numbers)

	/*
	   * 配列の要素数カウント
	     * 任意の配列を作成し、要素数をカウントし画面に表示する
	*/
	strings := [...]string{"hoge", "fuga", "foo", "bar"}
	pretty.Printf("#4 len(strings):%#v\n", len(strings))

}
