-- 看一下有没有被自己标记的消息，有说明上次没处理完，直接返回
local existMsg = redis.call('GET', KEYS[2])
if existMsg ~= false then
    return existMsg
end

while true do
    -- 读取消息列表左端的第一条消息
    local firstMsg = redis.call('LINDEX', KEYS[1], 0)
    if firstMsg == false then
        break
    end

    local isContinue = 0
    -- 将获取到的消息和其他订阅者标记的消息进行比较
    for index, value in ipairs(ARGV) do
        existMsg = redis.call('GET', value)
        -- 如果发现这条消息被其他订阅者标记了，就把这条消息从消息列表左端移除
        if firstMsg == existMsg then
            redis.call('LPOP', KEYS[1])
            isContinue = 1
            break
        end
    end

    -- 如果获取到的消息没有被其他订阅者标记，那就自己就标记，然后把这条消息从消息列表左端移除
    if 0 == isContinue then
        redis.call('SET', KEYS[2] ,firstMsg)
        redis.call('LPOP', KEYS[1])
        return firstMsg
    end
end

return ""
