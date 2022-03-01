package main

import (
	"github.com/yuin/gopher-lua"
)

// START OMIT
const scr = `
function print_t(t, idx)
  print("t[" .. tostring(idx) .. "]", t[idx])
end
function dump_t(t)
  if t[1] then
    for i,v in ipairs(t) do
      print("t[" .. tostring(i) .. "]", v)
    end
  else
    for k,v in pairs(t) do
      print("t[" .. tostring(k) .. "]", v)
    end
  end
end
-- luastart // OMIT
-- numeric keys
t_arr = { 3, 4, "5" }

-- indexing starts at 1 :-)))
dump_t(t_arr)

-- variable keys, can be accessed two different ways
k_str = "specific key"
k_tbl = {"table as key"}
t_hash = { 
  a = "a!", 
  b = true, 
  [k_str] = "specific value", -- [k] = "don't use k_str as the accessor"
  [k_tbl] = "table as key",
}

print("t.b", t_hash.b) -- "true"
print_t(t_hash, {"table as key"}) -- nil
dump_t(t_hash)
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
