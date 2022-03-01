function print_t(t, idx)
  print("t[" .. tostring(idx) .. "]", t[idx])
end
-- luastart // OMIT
-- numeric keys
t_arr = { 3, 4, "5" }

-- indexing starts at 1 :-)))
print_t(t_arr, 3) -- 5
print_t(t_arr, 1) -- 3

-- variable keys, can be accessed two different ways
k_str = "specific key"
k_tbl = {"table as key"}
t_hash = { a = "a!", b = true, k_str = "specific value", k_tbl = "table as key" }
print_t(t_hash, "a") -- "a!"
print("t.b", t_hash.b) -- "true"
print("t[k_str]", t_hash[k_str]) -- "specific value"
print("t[k_tbl]", t_hash[k_tbl]) -- "table key"
print("t[{\"table as key\"}]", t_hash[{"table as key"}]) -- nil


