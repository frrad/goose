package main

import (
	"fmt"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	prog, err := createProgram("examples/ex.go")
	if err != nil {
		panic(err)
	}

	fn := getFnByName(prog, "simplest")
	infer(fn)
}

func createProgram(path string) (*ssa.Program, error) {
	lconf := &loader.Config{}
	lconf.CreateFromFilenames("", path)

	lprog, err := lconf.Load()
	if err != nil {
		return nil, err
	}

	prog := ssautil.CreateProgram(lprog, 0)
	prog.Build()
	return prog, nil
}

func getFnByName(prog *ssa.Program, name string) *ssa.Function {
	fns := ssautil.AllFunctions(prog)

	for fn := range fns {
		if fn.Name() == name {
			return fn
		}
	}
	return nil
}

func infer(fn *ssa.Function) {
	fmt.Println(fn.Signature)
	fmt.Println(getRetVals(fn))

	for _, block := range fn.Blocks {
		fmt.Println("block #", block)
		fmt.Println("dominees", block.Dominees())
		for i, instr := range block.Instrs {
			fmt.Println(i, instr)
		}
		fmt.Println("")
	}
}

func getRetVals(fn *ssa.Function) ([]ssa.Value, error) {
	var ans []ssa.Value
	for _, block := range fn.Blocks {
		if ret, ok := containsReturn(block); ok {
			if ans == nil {
				ans = ret
			} else {
				return nil, fmt.Errorf(
					"can't handle more than one return in fn yet")
			}
		}
	}
	return ans, nil
}

func containsReturn(block *ssa.BasicBlock) ([]ssa.Value, bool) {
	for _, instr := range block.Instrs {
		if ret, ok := instr.(*ssa.Return); ok {
			return ret.Results, true
		}
	}
	return nil, false
}
