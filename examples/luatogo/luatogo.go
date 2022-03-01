package main

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
function add(a, b)
  return a+b
end
x = add(2, 3)
`

func main() {
	vm := lua.NewState()
	vm.DoString(scr)
	fn := vm.GetGlobal("add")
	fmt.Printf("add: %v\n", fn)
	x := vm.GetGlobal("x")
	fmt.Printf("x: %v\n", x)
	vm.CallByParam(lua.P{Fn: fn, NRet: 1}, lua.LNumber(1), lua.LNumber(2))
	y := vm.Get(1)
	fmt.Printf("y-from-go: %v", y)
}

// END OMIT
