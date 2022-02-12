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