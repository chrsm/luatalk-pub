local c = coroutine.create(function(a, b)
  print("args: " .. tostring(a) .. "," .. tostring(b))

  a, b = coroutine.yield(a+b, 0)
  print("back inside: " .. tostring(a) .. "," .. tostring(b))

  return "done"
end)

local ok, inita, initb = coroutine.resume(c, 1, 2)
print("initial: " .. tostring(inita) .. "," .. tostring(initb))
ok, inita, initb = coroutine.resume(c, 3, 4)
print("second: " .. tostring(inita) .. "," .. tostring(initb))
print("ret: " .. tostring(coroutine.status(c)))
print("typ: " .. tostring(type(c)))
