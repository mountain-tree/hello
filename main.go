package main

import (
	"example/hello/foo"
	"fmt"
	"path"

	"rsc.io/quote"
)

func main() {
	fmt.Println(foo.Bar())
	fmt.Println(quote.Go())

	var dir, file string

	dir, file = path.Split("~/.local")

	fmt.Println("dir:", dir)
	fmt.Println("file:", file)

	var f string
	_, f = path.Split("~/.local/")
	fmt.Println("file-f:", f)
	_, f = path.Split("~/.local")
	fmt.Println("file-f:", f)
}
