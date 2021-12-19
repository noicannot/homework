#Go进阶训练营第9周作业

## 1.总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。
## 2.实现一个从 socket connection 中解码出 goim 协议的解码器。


fix length
发送方，每次发送固定长度的数据，并且不超过缓冲区，接受方每次按固定长度区接受数据
delimiter based
发送方，在数据包添加特殊的分隔符，用来标记数据包边界
length field based
发送方，在消息数据包头添加包长度信息