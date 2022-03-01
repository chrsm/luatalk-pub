package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
example_var = 1.0
example_var = true
example_var = "assigning a new type is OK"

example_var = {
  seriously = "it is not",
  ["even remotely"] = "a problem",
}

function example_var(a, b)
  return a+b
end

example_var = example_var(1, 1)
print(example_var) -- output: 2
-- luaend // OMIT
`

// END OMIT

func main() {
	vm := lua.NewState()
	defer vm.Close()

	if err := vm.DoString(scr); err != nil {
		panic(err)
	}
}
