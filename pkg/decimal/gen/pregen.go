package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	var fh *os.File
	if fh, err = os.Create("base10k.txt"); chk.E(err) {
		panic(err)
	}
	for i := range 10000 {
		fmt.Fprintf(fh, "%04d", i)
	}
}
