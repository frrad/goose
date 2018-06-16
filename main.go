package main

import (
	"golang.org/x/tools/go/loader"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	lconf := &loader.Config{}
	lconf.CreateFromFilenames("./examples", "examples/ex.go")

	lprog, err := lconf.Load()
	if err != nil {
		panic(err)
	}

	mode := ssa.PrintPackages
	ssautil.CreateProgram(lprog, mode)
}
