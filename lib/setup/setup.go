package setup

import (
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func CreateProgram(path string) (*ssa.Program, error) {
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

func FnByName(prog *ssa.Program, name string) *ssa.Function {
	fns := ssautil.AllFunctions(prog)

	for fn := range fns {
		if fn.Name() == name {
			return fn
		}
	}
	return nil
}
