package main

import "fmt"

func AppendString(value1 string, value2 string) (value3 string) {
	/* Normal：文字列の連結 */
	value3 = value1 + value2
	return
}

func main() {
	fmt.Println(AppendString("ABC", "HOGE"))
}
