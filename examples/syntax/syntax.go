package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
print"hello!!"

oslib = require("os")

globalvar = "a variable decl"
if globalvar then
  print(globalvar)
end

function x(a, b)
  return a+b
end

tbl = { 1, 2, 3 } -- indexes start at 1: don't cry
for idx,v in ipairs(tbl) do
  print(tostring(idx) .. "=" .. v)
end

print("tbl len: " .. tostring(#tbl)) -- # is the length operator

local tbl_h = { a = 1, b = 2, ["c key"] = 3 }
for k,v in pairs(tbl_h) do
  print(k .. "=" .. v)
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
