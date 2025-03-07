-- 基于有序集合（sorted set）实现的滑动窗口限流算法

-- 限流对象
local key = KEYS[1]
-- 窗口总时长（毫秒）
local windowSize = tonumber(ARGV[1])
-- 最大请求数
local maxReqNum = tonumber(ARGV[2])
-- 当前时间（毫秒时间戳）
local now = tonumber(ARGV[3])

-- 窗口的起始时间
local windowStart = now - windowSize

-- '-inf' 表示第一个成员；'+inf' 表示最后一个成员；

-- 删除集合中所有 score <= windowStart 的成员
redis.call('ZREMRANGEBYSCORE', key, '-inf', windowStart)

-- 计算窗口总请求数
local count = redis.call('ZCOUNT', key, '-inf', '+inf')

if count >= maxReqNum then
    -- 被限流
    return "true"
else
    -- 没有限流

    -- score 和 member 都设置成 now 就行
    redis.call('ZADD', key, now, now)
    -- 更新限流对象的过期时间
    redis.call('PEXPIRE', key, windowSize)

    return "false"
end