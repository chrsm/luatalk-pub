package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
-- globally accessible
y = "2"
print("y\t", y)
_G.y = "3" -- special variable meaning "globals"
print("y (no _G)", y)

-- scoped to file
local x = "1"
print("x\t", x)
print("_G.x\t", _G.x)

do
  -- scoped to block
  local z = "3"
  print("z\t", z) -- output: 3
end

print("z (outside do)", z) -- output: nil
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
