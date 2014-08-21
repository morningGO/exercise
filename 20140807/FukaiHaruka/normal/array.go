/*
Morning Go...!! 第１期 2014年08月07日の宿題
Normal: 配列の操作
配列の連結 任意の配列を複数作成し、その配列を連結し画面に表示する
配列のスライス 任意の配列を作成し、配列をスライスし画面に表示する
配列値のソート 任意の配列を作成し、値に整数値を代入します。代入した値をソートしてまた配列に代入し画面に表示する
配列の要素数カウント 任意の配列を作成し、要素数をカウントし画面に表示する
*/

package main

import (
	"fmt"
	"unsafe"
)

// 配列の連結
func addArray(){
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	len_1 := len(arr1)
	len_2 := len(arr2)
	
	// 連結後の、配列length指定方法がわかりませんでした
	// length指定は定数式のみ有効なので、
	// var arr_result [(len_1 + len_2)]int のようなことができない
	var arr_result [10]int

	for i := 0; i < (len_1 + len_2); i++ {
		if i < len_1 {
			arr_result[i] = arr1[i]
		} else {
			arr_result[i] = arr2[i - len_1]
		}
	}
	
	fmt.Printf("=== Homework Normal 1: AddArray ===\n")
	fmt.Printf("arr1 = %v\n", arr1)
	fmt.Printf("arr2 = %v\n", arr2)
	fmt.Printf("Add arr1 + arr2 result = %v\n\n", arr_result)
}

// 配列のスライス
func sliceArray(){
	arr_org := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr_org[0:len(arr_org)]
	
	fmt.Printf("=== Homework Normal 2: SliceArray ===\n")
	fmt.Printf("arr_org = %v\n", arr_org)
	fmt.Printf("Slice arr_org[0:len(arr_org)] result = %v\n\n", slice)
}

// 配列のソート(バブルソート)
func sortArray(){
	arr_org := [6]int{4, 2, 1, 6, 5, 3}
	var arr_result [6]int
	arr_result = arr_org

	for i := 0; i < (len(arr_result) - 1); i++ {
		for j := (len(arr_result) - 1); j > i ; j-- {
			if arr_result[j - 1] > arr_result[j] {
				temp := arr_result[j]
				arr_result[j] = arr_result[j - 1]
				arr_result[j - 1] = temp
			}
		}
	}
	fmt.Printf("=== Homework Normal 3: SortArray ===\n")
	fmt.Printf("arr_org = %v\n", arr_org)
	fmt.Printf("Sort result = %v\n\n", arr_result)
}

//配列の要素数カウント
func countArray(){
	arr_org := [6]int{1, 2, 3, 4, 5, 6}
	count := unsafe.Sizeof(arr_org) / unsafe.Sizeof(arr_org[0])
	fmt.Printf("=== Homework Normal 4: CountArray ===\n")
	fmt.Printf("arr_org = %v\n", arr_org)
	fmt.Printf("Count result = %d\n\n", count)	
}

func main(){
	// 配列の連結
	addArray()
	
	// 配列のスライス
	sliceArray()

	// 配列のソート
	sortArray()
	
	//配列の要素数カウント
	countArray()
}