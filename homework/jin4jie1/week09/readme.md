### 作业1

> 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

#### fix length，定长编码

定长编码粘包：因为编码后报文长度是固定的，每次从缓冲区，按定长取即可

定长编码应用：比如平常解析文本的时候，如果把字符缓冲区按位再分的细一点，可以看做是在解析定长的 ascii 码，不过底层帮助屏蔽掉了

#### delimiter based，基于定界符

基于定界符粘包：在缓冲区找到界定符，就可以分隔两个包了。

基于定界符应用：http 1.1 的从实操来看就可以算是界定符类型的。找到 \r\n\r\n 就表示 http 1.1 的报文头部已经完整了，再在报文头部找到 content length 就可以计算出整个包的长度。

#### length field based frame decoder，基于长度解码

基于长度解码粘包：主要是需要根据特征位先判断用几个字节表示长度，然后就可以计算出报文长度。

基于长度解码应用：WebSocket 感觉可以算是一种基于长度解码的。WebSocket 的 8 bit 的 Payload len 字段：如果小于 126 就直接表示报文长度；如果等于 126 那么后面的 2 字节就表示报文长度；如果等于 127 那么后面的 8 字节就表示报文长度。

### 作业2

> 实现一个从 socket connection 中解码出 goim 协议的解码器。

- 目录下的 service/main.go 是一个 tcp 服务端，负责接收 client 发来的 goim 报文然后解码
- 目录下的 client/main.go 是一个 tcp 客户端，负责编码 goim 报文然后发送给 service

客户端和服务端都是一次性的，简单的实现编码并发送和接收并解码报文的逻辑。

之前写过 http 1.1 和 websocket 的，报文解析的逻辑看上去都差不多。就在这个作业目录的隔壁：

[https://github.com/KelipuTe/demo_golang/tree/master/tcp_service_v1](https://github.com/KelipuTe/demo_golang/tree/master/tcp_service_v1)

实现了 http 1.1，自定义字节流，websocket：

[https://github.com/KelipuTe/demo_golang/tree/master/tcp_service_v1/internal/protocol](https://github.com/KelipuTe/demo_golang/tree/master/tcp_service_v1/internal/protocol)