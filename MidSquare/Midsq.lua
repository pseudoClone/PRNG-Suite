seed = tostring(os.time())
seed = seed:sub(-4) -- Apparently you need modules to handle bigint
print(seed)

function MidSq()
    local square = tostring(seed * seed)
    while #square ~= 12 do  -- #s => len(s)
        square = "0" .. square -- Padding
    end
    seed = tonumber(square:sub(4, 9)) --Slice mid 4
    return seed
end

while true do
    print(MidSq())
end
