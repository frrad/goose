package main

import (
	"fmt"

	"golang.org/x/tools/go/ssa"

	"github.com/frrad/goose/lib/setup"
)

func main() {
	prog, err := setup.CreateProgram("examples/ex.go")
	if err != nil {
		panic(err)
	}

	fn := setup.FnByName(prog, "dagFn")
	infer(fn)
}

func infer(fn *ssa.Function) {
	if isDAG(fn) {
		inferDAG(fn)
		return
	}

	panic("oh no!")
}

func inferDAG(fn *ssa.Function) {
	startingBlock := fn.Blocks[0]
	fmt.Println(startingBlock)
}

//isDAG checks if the graph for a function is a DAG
func isDAG(fn *ssa.Function) bool {
	// todo: implement this function
	return true
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
