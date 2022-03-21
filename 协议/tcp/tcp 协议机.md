# TCP协议机
本身是不太想来把这部分归结到这里的，但方便后面自己查阅的时候快速略过。协议机主要将建立tcp的时候参与数据传输的双方的
状态，最开始想把这个归结到linux上面，但是想了一下放在linux下可能只能是非常小的一部分，甚至在观看的时候会被漏掉，如果
放在这里的话至少可以和tcp协议一起看了。

### 一个完整的数据传输过程是怎么样的
- 客户端： syn , win , mss 
- 服务器： ack ，syn，win，mss
- 客户端：ack ，win，mss

- 服务器： data1 
- 服务器： data2
- 客户端： ack data1 、data2 

- 客户端： fin 
- 服务器： ack 
- 服务器： ack ，fin 
- 客户端： ack


上面就是一个完整的传输过程，接下来我们看看客户端、服务器双方的一个状态


### 客户端

- 1.发送syn ，进入syn_sent 
- 2.接受到 ack 、 syn ，并发送ack ， 进入establish
- 。。。。 数据沟通都是establish 。。。。
- 3.发送fin ，进入fin_wait_1 状态
- 4.接受到服务器的ack ，进入fin_wait_2 状态
- 5.接受到 fin，ack ，发送ack ，进入time_wait 
- 6.等待30s ，close

客户端总状态： syn_sent,fin_wait_1,fin_wait_2,time_wait 
### 服务器
- 1.接收到syn 包，并回复syn，ack ，此时进入syn_rcvd 
- 2.接收到ack 包，进入establish
- 。。。。 中间过程是establish ....
- 3.收到了 fin ，并回应ack， 进入close_wait
- 4.发送 fin，ack ，进入last_ack
- 5.接受ack ，进入close状态

服务器总结： syn_rcvd、 close_wait、last_ack