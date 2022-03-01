package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
t_arr = { 3, 4, "5" }

table.remove(t_arr, 2)

for i,v in ipairs(t_arr) do
  print("t_arr[" .. tostring(i) .. "]", v)
end

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
