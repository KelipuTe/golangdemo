local existMsg = redis.call('GET', KEYS[2])
if existMsg == false then
    return ""
end

redis.call('RPUSH', KEYS[1], existMsg)

redis.call('DEL', KEYS[2])

return ""