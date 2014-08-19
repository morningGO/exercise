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
	if err := YAMLToFiles("../gopher_files.yml"); err != nil {
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
