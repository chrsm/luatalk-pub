package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
local yield, resume, status = 
      coroutine.yield, coroutine.resume, coroutine.status

local coro = coroutine.create(function(initial, max)
  local v = initial
  while v < max do
    yield(v)
	v = v + 1
  end

  return v -- get to max
end)

while status(coro) ~= "dead" do
  ok, v = resume(coro, 10, 15)
  print("v =", v, ok, status(coro))
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
