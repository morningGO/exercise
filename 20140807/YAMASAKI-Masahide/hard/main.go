/*
## Hard:

* GNU core utilities のコマンドの一つであるcatコマンドを実装してください。
* http://www.gnu.org/software/coreutils/
* すべての機能を実装しなくていいです。catコマンドのメイン機能である2つ以上のファイルを連結して標準出力に表示する。
※もし余裕があれば、WindowsとMacOSXとLinuxに対応したソースコードを書いてみてください。余裕があれば。
*/

package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(os.Stdout, file)
	}
}
