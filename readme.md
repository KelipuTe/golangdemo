## 说明（readme）

放 golang 代码，一些聚焦于单个知识点的示例，一些自己实现的框架。

### 目录里面是什么

- "algorithm/"，算法。
- "cache/"，自己实现的缓存框架。
- "demo/"，聚焦于单个知识点的示例。
- "leetcode/"，leetcode 算法题。
- "orm/"，自己实现的 orm 框架。
- "rpc/"，自己实现的 rpc 协议。
- "web/"，自己实现的 web 框架。

缓存框架。

- 主要实现：
    - 缓存框架
    - 缓存框架支持本地缓存作为底层
    - 缓存框架支持 redis 作为底层
    - redis 锁
- 次要实现：
    - 内存限制和 lru 淘汰策略
    - read through
- 计划实现：
    - 其他的缓存模式

自己实现的 rpc 协议。

- 主要实现：
    - 自定义 rpc 协议
    - rpc 客户端
    - rpc 服务端
- 次要实现：
    - 自定义 json 协议
    - json 序列化
- 计划实现：
    - 增加 protobuf 序列化
    - 增加 gzip 压缩
