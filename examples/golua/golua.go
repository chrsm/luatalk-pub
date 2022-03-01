package main

// START OMIT
import (
	"github.com/yuin/gopher-lua"
)

const scr = `
print("hi!")
`

func main() {
	vm := lua.NewState()
	defer vm.Close()

	if err := vm.DoString(scr); err != nil {
		panic(err)
	}
}

//END OMIT
