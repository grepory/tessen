package main

import (
	"os"

	"github.com/grepory/tessen/check"
)

func main() {
	check.NewCheck(
		"HTTPCheck",
		HTTPInquisitor{
			Target: os.Args[1],
		},
	).Run()
}
