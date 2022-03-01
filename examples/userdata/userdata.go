package main

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func main() {
	vm := lua.NewState()
	defer vm.Close()

	setupVM(vm)

	if err := vm.DoString(example_init); err != nil {
		panic(err)
	}
}

//START OMIT
type Data_t struct{ V int }

func (d *Data_t) DoThing(a, b int) {
	d.V = a + b
}

func setupVM(ls *lua.LState) {
	data_t := new(Data_t)
	ls.SetGlobal("data", luar.New(ls, data_t))
}

const example_init = `
data:DoThing(1, 2)
print(data.V)
`

//END OMIT
