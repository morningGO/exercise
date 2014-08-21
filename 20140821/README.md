# MorningGo...!! 第1期 2014年08月21日の宿題

## Normal：ファイルの読み書き

1. ファイルに任意の文字列を書き出す
2. 1.で書き出したファイルを読み出し、書き出した文字列と比較し等しいことを確認する

## Hard: 謎のYAMLファイルを解読

以下の`main_test.go`のテストコードが通るように、`main.go`の`YAMLToFiles`関数を実装してください。
([/20140821/YAMASAKI-Masahide/hard](https://github.com/morningGO/exercise/tree/master/20140821/YAMASAKI-Masahide/hard) を各自のディレクトリにコピーしてください)

### テストの実行方法

```bash: 
$ cd 20140821
$ cp -a YAMASAKI-Masahide <ユーザー名>  # テンプレートをコピー
$ cd <ユーザー名>/hard                  # ディレクトリに移動 
$ go test                               # テスト実行
--- FAIL: TestYAMLToFiles (0.00 seconds)
	main_test.go:41: Not implemented
FAIL
exit status 1
FAIL	github.com/morningGO/exercise/20140821/YAMASAKI-Masahide/hard	0.008s
### ↑失敗するとこのようなエラーが出る ##

### 成功すると↓ ########################
$ go test
PASS
ok  	github.com/morningGO/exercise/20140821/YAMASAKI-Masahide/hard	0.009s
```

*テンプレートコード*

* main.go

```go
package main

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
)

type File struct {
	Name   string
	Base64 string
}
type Dir struct {
	Name  string
	Files []File
}

func main() {
	if err := YAMLToFiles("../../gopher_files.yml"); err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		indexFile := filepath.Join("gopher_files", "index.html")
		switch runtime.GOOS {
		case "windows":
			exec.Command("cmd", "/c", "start", indexFile).Start()
		case "darwin":
			exec.Command("open", indexFile).Start()
		default:
			fmt.Errorf("ブラウザで以下のファイルを開いて下さい。\n%s\n", indexFile)
		}
	}
}

func YAMLToFiles(filename string) error {

	//TODO: ここに実装する
	return errors.New("Not implemented")

}
```

* main_test.go

```go
package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testYAML = `
name: %s
files:
  - name: fibonacci.go
    base64: H4sIAAAJbogA/2SPsW7DIBRFZ7+vuMpkq47TrpWyZu4vPCgkqIAjeGSp8u99xM3SLrZA5557ubL94rND4pCJQrquRbDzSXZEhwN8MChOWskVDN+ylbBmyIXled+x2qx1tYabwymYNbO1Abkl40pdqMe6aZweBv2FLPimgWcYvB/xOuONhs33B9mYI8wMfjF6/qWYhjvdaXP38RpR3Hfdo4oG3XV6DrYco76gOLgbx8biPhGdl72s+xLOF1k0m2T5KNob8+jHacb/z6SVPwEAAP//QqPtTTIBAAA=
  - name: helo.go
    base64: H4sIAAAJbogA/xzKMQ6CQBBG4do5xe9WkBg4hI2lreVkGWDCMkuWpTDGG9jqWTyQiddwY/u+17a4xA2eDdJpRh51hY+d7KnQMaifMEoSsHVYM6eyXBe1oaGF/cSDYGY1Ip2XWND1c3ZE/Wb+D1WNG+1KbM5JLQer3ElCiAd83q/v8+FqutMvAAD//8MixFaFAAAA
`

var testPattern = []struct {
	filename string
	sha1     string
}{
	{"fibonacci.go", "9fc1fd1d9d613190a61c61bd842b775b72a2c99c"},
	{"helo.go", "eabca490026b91f0587aa7e801e2ca716d7a674b"},
}

func TestYAMLToFiles(t *testing.T) {

	tmpOutputDirName := "tmpdir-" + makeRandomString()
	tmpYAMLfileName := makeTmpYAML(tmpOutputDirName)

	err := YAMLToFiles(tmpYAMLfileName)
	defer os.RemoveAll(tmpYAMLfileName)
	defer os.RemoveAll(tmpOutputDirName)
	if err != nil {
		t.Fatal(err)
	}

	for _, tp := range testPattern {
		s, err := sha1sum(filepath.Join(tmpOutputDirName, tp.filename))
		if err != nil {
			t.Error(err)
		}
		if s != tp.sha1 {
			t.Errorf("sha1(%q) = <%s> want <%s>", tp.filename, s, tp.sha1)
		}

	}

}

func sha1sum(filename string) (string, error) {
	h := sha1.New()
	fio, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	io.Copy(h, fio)
	fio.Close()
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func makeTmpYAML(dirname string) (filename string) {
	tmpYAMLfile, _ := ioutil.TempFile(`.`, `test`)
	tmpYAMLfile.Write([]byte(fmt.Sprintf(testYAML, dirname)))
	filename = tmpYAMLfile.Name()
	tmpYAMLfile.Close()
	return
}

func makeRandomString() string {
	rb := make([]byte, 18)
	rand.Read(rb)
	return base64.URLEncoding.EncodeToString(rb)
}
```


## 提出者





