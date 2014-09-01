package main

import (
    "os"
    "io"
    "fmt"
)

func readWrite(src io.Reader, dst io.Writer) {
}

func main() {
    if len(os.Args) == 1 {
        _, err := io.Copy(os.Stdout, os.Stdin)
        if err != nil {
        fmt.Println(err)
        }
    } else {
        for _,fname := range os.Args[1:] {
            fh, err := os.Open(fname)
            if err != nil {
                fmt.Println(err)
            }

            _, err = io.Copy(os.Stdout, fh)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}



