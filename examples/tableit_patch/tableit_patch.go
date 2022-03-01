package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
local _next = next
function next(t, i)
  local mt = getmetatable(t)
  local fn = mt and mt.__next or _next

  return fn(t, i)
end
-- luastart // OMIT
t_arr = { 3, nil, 4, "5" }
setmetatable(t_arr, { __index = function(t, i)
  local ni = next(t, i)
  if ni == nil then
    -- try i+1
	ni = next(t, i+1)
  end

  return ni
end })

t_arr[1] = nil
for i,v in next, t do
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
