package main

// START OMIT
import (
	"github.com/yuin/gopher-lua"
)

const scr = `
-- luastart // OMIT
function add(a, b)
  return a+b
end

function run(fn, ...)
  return fn(...)
end

print(run(add, 1, 2))

run(function(a, b)
  print(a, b)
end, 1, 2)
-- luaend // OMIT
`

func main() {
	vm := lua.NewState()
	defer vm.Close()

	if err := vm.DoString(scr); err != nil {
		panic(err)
	}
}

//END OMIT
