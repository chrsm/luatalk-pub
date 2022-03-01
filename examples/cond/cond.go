package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
if a then
  print("a exists!")
elseif not a then
  print("a doesn't exist!")
end

b = 2
if b ~= 1 then
  print("b != 1")
end

a = 10
while a > 0 do
  a = a-1
  if a % 2 == 0 then print("a = ", a) end
end

b = 10
repeat
  b = b - 1
  if b % 2 == 0 then print("b = ", b) end
until b == 0
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
