# String 
在redis中的数据结构中说万精油的，该说不说就是string了，那么今天就来看看string的数据结构到底是怎么回事。内存耗费是怎么回事
内部内存布局是怎么回事，redis对string做了哪些优化。

## 目录
- [string内存开销](##string内存开销)
    - [string开销](####string开销)
    - [redisObject开销](####redisObject开销)
    - [dictEntry开销](####dictEntry开销)
    - [额外的内存开销](####额外的内存开销)
    - [简单的总结](####简单的总结)
- [string的优化](##string的优化)
    - [long类型优化](####long类型优化)
    - [embstr优化](####embstr优化)
    - [raw优化](####raw优化)
    
    
### string的内存开销
首先我们要确认的是什么？ 那么就是string的SDS 结构体，sds 内部存在3个内容
- len :4B，真是数据长度
- buf: 字节数组 以\0结尾。
- alloc ：类似golang数组的cap 大小4B

那么这里开销就有：除了真实数据外还有大概 8B的额外开销，但是这里仅仅是redis的string的基础结构开销。如果我需要寻址和一些额外属性
该怎么办？ 就要向上分析了，这里只是最底层

#### RedisObject 的内存开销
RedisObject 内部的内容是什么呢？
- 元数据 一个指针  8B
- 真实数据 一个指针指向 sds 8B

那么这里我们就可以知道了 存一个string的额外开销在redisobj上面就是16字节。在这里总的额外开销是24字节。

#### dictEntry 的内存开销
- *key 指向key的指针 8B。 这里要注意一下key的类型也是redisObject
- *Value 指向value 的指针 8B ，value的类型也是redisObject
- *Next 指向next dictEntry 的指针 8B

那么这里单纯的一个全局key 的内容大概就是24字节。总的开销再加上上面的开销：48字节，你可能会说48字节还不够我存一句话的，那我举个简单
例子来说把： 这个可以存24个中文字符，可以存48个英文字符。你看是不是这个开销已经超过很多时候内容本身的存储开销了 

#### 额外的内存开销
为什么还有其他额外的开销呢，和redis的 内存分配是挂钩的，redis内存是jemalloc来分配的，这个分配的话为了减少频繁的分配带来的性能开销
jemalloc 总是分配到2的次方的内存，比如我们开销为30字节，那么就会分配到32 字节，就会多2个字节的开销。
#### 简单的总结
- 一个rediskey 开销24字节
- 一个redisObject 开销是16字节
- 那么一个rediskey 开销为40 字节。

### string的优化
redis 在string上面做了很多优化来节省空间以及减少内存的碎片，主要有下面几种优化模式


#### long类型优化
当保存的内容是整数的时候，就会自动使用long类型来进行优化，具体优化存储方式就不再使用sds了，而是直接在
redisObject 上面存储的指针换成longint类型进行存储，就不需要sds的开销了

#### embstr优化
当字符串长度小于等于44字节的时候，redisObject是放在一起的，减少内存的碎片化，类似数组一样的内存布局
还可以利用cpu预读，可以提升性能。
#### raw优化
当字符串长度大于44字节，redisObject是分开的，类似链表的一样的内存布局