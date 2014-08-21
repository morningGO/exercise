/*
Morning Go...!! 第１期 2014年08月07日の宿題
Hard: catコマンドの実装
すべての機能を実装しなくていい。
catコマンドのメイン機能である2つ以上のファイルを連結して標準出力に表示する。
※もし余裕があれば、WindowsとMacOSXとLinuxに対応したソースコードを書いてみてください。→余裕ありませんでした……
*/

package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main(){	
	var arr1, arr2 []byte
	arr1, _ = ioutil.ReadFile(os.Args[1])
	arr2, _ = ioutil.ReadFile(os.Args[2])
	arr_res := append(arr1, arr2...)
		
	for i:=0; i < len(arr_res); i++ {
        fmt.Print(string(arr_res[i]))
    }
}
