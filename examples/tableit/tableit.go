package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
t_arr = { 3, 4, "5" }
t_hash = { a = "a!", b = true }

-- numeric keys iterated with ipairs
for i,v in ipairs(t_arr) do
  print("t_arr[" .. tostring(i) .. "]", v)
end

-- variable keys with pairs
for k,v in pairs(t_hash) do
  print("t_hash[" .. tostring(k) .. "]", v)
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
