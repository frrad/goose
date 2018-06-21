package main

import (
	"fmt"

	"golang.org/x/tools/go/ssa"

	"github.com/frrad/goose/lib/properties"
	"github.com/frrad/goose/lib/setup"
)

func main() {
	prog, err := setup.CreateProgram("examples/ex.go")
	if err != nil {
		panic(err)
	}

	// fn := setup.FnByName(prog, "simplest")
	fn := setup.FnByName(prog, "dagFn")
	infer(fn)
}

func infer(fn *ssa.Function) {
	showFn(fn)
	if isDAG(fn) {
		inferDAG(fn)
		return
	}

	panic("oh no!")
}

func showFn(fn *ssa.Function) {
	for i, bl := range fn.Blocks {
		fmt.Println("\nblock", i)
		for j, instr := range bl.Instrs {
			fmt.Println(j, ":", instr)
		}
	}
	fmt.Printf("\n")
}

func inferDAG(fn *ssa.Function) properties.Prop {
	startingBlock := fn.Blocks[0]
	instr := startingBlock.Instrs[1]

	switch in := instr.(type) {
	case *ssa.BinOp:
		fmt.Println("it's a binop")
		fmt.Println(in.X)
		fmt.Println(in.Y)
	case *ssa.Return:
		fmt.Println("it's a ret")
		fmt.Println("results:", in.Results)
	case *ssa.Phi:
		fmt.Println("it's a phi function")

	default:
		fmt.Printf("I don't know about type %T!\n", in)
	}

	return properties.Prop{}
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
