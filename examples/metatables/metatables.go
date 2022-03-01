package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
local vec2 = {}
vec2.__index = vec

function vec2.new(x, y)
  local self = setmetatable({}, vec2)
  self.x, self.y = x, y
  return self
end

vec2.__eq = function (a, b)
  return a.x == b.x and a.y == b.y
end

local v1 = vec2.new(1, 1)
local v2 = vec2.new(2, 2)

print(tostring(v1 == v2)) -- output: false

v2.x, v2.y = 1, 1
print(tostring(v1 == v2)) -- output: true

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
