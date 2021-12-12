#Go进阶训练营第8周作业


##1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
###命令:
```
redis-benchmark -d 10 -t get,set
```
###SET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 0.89 seconds|112866.82 requests per second|
|20|100000 requests completed in 0.89 seconds|112233.45 requests per second|
|50|100000 requests completed in 0.87 seconds|114942.53 requests per second|
|100|100000 requests completed in 0.86 seconds|116009.28 requests per second|
|200|100000 requests completed in 0.86 seconds|115874.86 requests per second|
|1k|100000 requests completed in 0.87 seconds|114547.53 requests per second|
|5k|100000 requests completed in 0.93 seconds|107526.88 requests per second|
###GET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 0.89 seconds|112485.94 requests per second|
|20|100000 requests completed in 0.85 seconds|117370.89 requests per second|
|50|100000 requests completed in 0.84 seconds|118906.06 requests per second|
|100|100000 requests completed in 0.86 seconds|116414.43 requests per second|
|200|100000 requests completed in 0.85 seconds|117785.63 requests per second|
|1k|100000 requests completed in 0.88 seconds|114285.71 requests per second|
|5k|100000 requests completed in 0.92 seconds|108225.10 requests per second|

##2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
###结果在csv文件

