package main

import (
	"fmt"
	"os"

	"github.com/mazei513/slice-gen/generator"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: github.com/mazei513/slice-gen {Name} {ObjType}")
		fmt.Println(" example:")
		fmt.Println(" github.com/mazei513/slice-gen User user'")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if err := generator.Generate(os.Args[1], os.Args[2], wd); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
