package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	jsonObjects := readFile(900374, 900375, "files")
	spew.Dump(jsonObjects)
}
