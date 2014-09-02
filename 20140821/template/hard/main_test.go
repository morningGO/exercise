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
