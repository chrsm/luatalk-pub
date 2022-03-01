local vec = {}
vec.__index = vec

function vec.new(x, y)
  local self = setmetatable({}, vec)
  self.x, self.y = x, y
  return self
end

vec.__eq = function (a, b)
  return a.x == b.x and a.y == b.y
end

local v1 = vec.new(1, 1)
local v2 = vec.new(2, 2)

print(tostring(v1 == v2)) -- output: false

v2.x, v2.y = 1, 1
print(tostring(v1 == v2)) -- output: true

