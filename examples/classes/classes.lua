local Player = {}
Player.__index = Player

function Player.new(name)
  local self = setmetatable({}, Player)
  self.name = name
  self.health = 100
  return self
end

function Player:hit(dmg) self.health = self.health - dmg end
function Player:alive() return self.health > 0 end

-- create a new Player
local p = Player.new("chrsm")
p:hit(150)
if not p:alive() then
  print(p.name .. " is dead!")
end -- output: chrsm is dead!
