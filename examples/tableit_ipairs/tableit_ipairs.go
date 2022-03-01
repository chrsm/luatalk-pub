package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
t_arr = { 3, nil, 4, "5" }

for i,v in ipairs(t_arr) do
  print("t_arr[" .. tostring(i) .. "]", v)
end

-- if we remove [1], nothing prints
t_arr[1] = nil
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
