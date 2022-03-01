package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
-- for init,expr do
-- note: init is explicitly local!
for i=1,10 do
  if i % 2 == 0 then
    print("i =", i)
  end
end
print("i =", i)

-- for init,expr,postexpr
for j=1,10,2 do
  print("j =", j)
end

-- break
for k=1,10 do
  print("k =", k)
  break
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
