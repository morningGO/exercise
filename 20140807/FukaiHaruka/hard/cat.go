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
	"flag"
	"bufio"
	"fmt"
)

func main(){	
	var arr1, arr2 []byte
	var err error
	var fp1, fp2 *os.File

	fp1, err = os.Open(flag.Arg(0), os.O_RDONLY, 0660)
	reader := bufio.NewReader(fp1)
	reader.Read(arr1)

	fp2, err = os.Open(flag.Arg(1), os.O_RDONLY, 0660)
	reader = bufio.NewReader(fp2)
	reader.Read(arr2)
	
	fmt.Printf("%v\n", append(arr1, arr2...))
}
