#Hash 
hash表真的是用hash存储数据的吗，这个点你能理解吗？



## 目录
- [底层实现](##底层实现)
    - [ziplist](##ziplist)
    
    
    
### ziplist
 在ziplist 那篇文章中已经讲解过了底层实现了，这次就是看一下ziplist在哈希表的实际应用。
 
 - hash-max-ziplist-entries : 表示用压缩列表保存哈希集合中的最大元素的个数
 - hash-max-ziplist-value : 表示用压缩列表保存哈希的单个元素最大长度
 
 如果我们往哈希中写入超过 hash-max-ziplist-entries的个数或者单个元素值超过hash-max-ziplist-value
 redis就会把hash的结构转化成为哈希表。否则就是以ziplist的方式进行存储的。
 
 同理也就是说redis如果在hash-max-ziplist-entries个数内是利用ziplist进行存储的，那么时间复杂度其实是
 On的一个时间复杂度，但是空间复杂度就相对较小。
 
 
 #### hash的扩容大小
  当下面的任意一个条件被满足的时候，hash表就会进行扩展，hash表的扩容倍数是2倍。虽然大家说的都是两倍，但是实际
  代码并不是两倍：这里的伪代码我用golang编写了一下
  
 ````
   func newPower(num)int {
      var i:= 0
      for num< 2^i{
             i ++ 
         }
         return 2^i
     }
 ````
 在这里你可以看到，其实扩容实际就是找到最接近分配值的内容，比如，我本次size=5 我不够用了，需要扩容到5*2 =10 ，
 此时将10 传入这个函数，算出最接近的内容是8，所以分配值为8，而不是10。具体扩容条件：
 
  - 1.服务器目前没有在执行任何的bgsave 或者bgrewriteAOF命令，且hash的负载因子大于等于1
  - 2.服务器目前正在执行bgsave或者bgrewriteAOF命令，且hash负载因子大于等于5
  - 负载因子计算方式： load_factor = ht[0].used/ht[0].size
  
这里你大概可以知道，扩容会发生的情况是hash表满了，而且没有在进行备份和重写，就会扩容redis，但是第二条如何去理解
呢？ 
 
 ### hash的扩容和收缩

 

 