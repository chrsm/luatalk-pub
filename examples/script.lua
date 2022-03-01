local _module_0 = { }
local p, to_lua
do
	local _obj_0 = require("yue")
	p, to_lua = _obj_0.p, _obj_0.to_lua
end
local unpackstr
unpackstr = function(fp)
	local n = unpack("i8", fp:read(8))
	return unpack("c" .. tostring(n), fp:read(n))
end
local stringFromFile
do
	local _with_0 = fp
	if _with_0 ~= nil then
		unpackstr(fp)
	end
	stringFromFile = _with_0
end
if not stringFromFile then
	error("damn it isn't there")
end
_module_0[#_module_0 + 1] = unpackstr
return _module_0
