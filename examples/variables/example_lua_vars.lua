example_var = 1.0

example_var = true

example_var = "assigning a new type is OK"

example_var = {
  seriously = "it is not",
  ["even remotely"] = "a problem",
}

function example_var(a, b)
  return a+b
end

example_var = example_var(1, 1)
print(tostring(example_var))

-- output: 2

