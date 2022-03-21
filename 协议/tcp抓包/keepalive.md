# KeepAlive
关于keepalive 是什么我也不想过多赘述，做过一些开发的人员都是比较懂的，但是keepalive 在tcp
上的具体作用是什么，以及具体keepalive的包是由那端发出，以及包的表现形式是怎么样的，还是希望
你有个具体的概念。


### 非默认配置
keep alive 在tcp的应用都不是默认配置，如果你要启动keep alive 的话就可以使用setsocketopt（）
系统调用，对以及创建的socket 进行配置进行配置，在你创建tcp服务器的时候加一个参数应该就可以了。
相关配置：

1. 间隔时间：net.ipv4.tcp_keepalive ,默认值为7200秒
2. 最大探测次数： net.ipv4.tcp_keepalive_probes ,在探测没有响应的话可以发出多次探测
3. 最长间隔时间： net.ipv4.tcp_keepalive_intvl ,多次探测之间最大间隔时间

其实你可以看的出来tcp 的keepalive 还是比较反应较慢的一个配置，一般来说为了减少网络开销，提高
系统的灵敏度，还是比较建议自己在应用层维持心跳


### keepalive的结构
在keepalive的结构报文中

1. 整体的交互流程是和握手挥手类似
2. 虽然心跳包len=0 ，但是当前心跳包是上一个包seq -1 