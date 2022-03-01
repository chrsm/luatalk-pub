package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
------- strings
hello = "world" -- immutable

-- can be concatenated!
print("hello " .. hello)

-- can use double or single quotes and include escape sequences
print('hello\tworld')

-- yuck: can be auto-converted to numbers
print(tostring("10" + 1))

------- numbers: internally, all doubles
n = 1
n = 3.001

-- can be cast to a string as seen above
print(tostring(n))

------- bools
b = true
if b then
  print("b = " .. tostring(b))
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
