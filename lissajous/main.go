package main

import (
	"lissajous/lj"
	"os"
)

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lj.Lissajous(os.Stdout, 5)
}
