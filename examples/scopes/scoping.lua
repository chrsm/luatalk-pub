-- scoped to file
local x = "1"
-- globally accessible
y = "2"

do
  -- scoped to block
  local z = "3"
  print(z) -- output: 3
end

print(z) -- output: nil
