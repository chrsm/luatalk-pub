local somelib = require("lua_file_no_extension")

globalvar = "writes to _G!"
local localvar = "stays in this scope unless returned!"

function x(a, b)
  return a+b
end

local tbl = { 1, 2, 3 } -- indexes start at 1: don't cry
for idx,v in ipairs(tbl) do
  print(tostring(idx) .. "=" .. v)
end

local tbl_h = { a = 1, b = 2, ["c key"] = 3 }
for k,v in pairs(tbl_h) do
  print(k .. "=" .. v)
end

