package main

import (
	"reflect"

	lua "github.com/yuin/gopher-lua"
	//	luar "layeh.com/gopher-luar"
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
type WindowManager struct{}

func setupVM(ls *lua.LState) {
	ghetto := new(WindowManager)

	userdata := ls.NewUserData()
	userdata.Value = reflect.ValueOf(ghetto).Interface()

	ls.SetGlobal("wm", userdata)
}

const example_init = `
print(wm)
`

//END OMIT
