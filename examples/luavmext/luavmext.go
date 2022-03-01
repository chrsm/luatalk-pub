package windows

import (
	user "github.com/chrsm/winapi/user"
	lua "github.com/yuin/gopher-lua"
)

//START OMIT
var exports = map[string]lua.LGFunction{
	"get_active_window": getActiveWindow,
}

func Preload(ls *lua.LState) {
	ls.PreloadModule("windows", func(ls *lua.LState) int {
		mod := ls.SetFuncs(ls.NewTable(), exports)
		ls.Push(mod)
		return 1
	})
}

func getActiveWindow(ls *lua.LState) int {
	hwnd := uintptr(user.GetActiveWindow())
	ls.Push(lua.LNumber(hwnd))
	return 1
}

//END OMIT
