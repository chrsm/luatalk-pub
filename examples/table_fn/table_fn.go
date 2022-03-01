package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
-- luastart // OMIT
t_thing = { 
  myval = 1,
  add = function(self, v)
    return self.myval + v
  end,
}

print(t_thing:add(2)) -- 3

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
