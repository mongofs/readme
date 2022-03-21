// 在redis 中，这个结构是可以在dict.h 中可以看到的
package dict

// 字典：dict
type dict struct {
	// dictType 是dict进行操作的集合，在使用go编写代码的时候喜欢使用结构方法，但是如果是包裹
	// 一个具体的interface的时候可以将具体的方法和结构进行解耦，这是属于一种编程方式
	dictType DictType
	// 这里保存的是dict的一些私有参数，在调用dictType的时候需要讲privdata整个进行传入进去
	privdata *privdata
	// 实际存储数据的两个hash表，通常使用的是hash[0]，只有在进行rehash的时候才会使用hash[1]
	dictht [2]*Dictht

	// 在使用rehash的时候，进行标记rehash的进程，如果在查询key=demo的时候，正在进行rehahsh，那么
	// hash(demo)< trehashidx,的时候，就直接在hash一遍去ht【1】中查找
	trehashidx int
}


// 这个有兴趣的可以去看一下
type privdata struct {
}


// dict字典底层存放数据的hash
type  Dictht struct {

	// 具体存放元素的数组，注意，这里是数组来的
	table [] *DictEntry

	// 这是hash表声明的大小
	size int64

	// 这是hash表进行计算下标计算的
	sizemask int64

	// 这个字段是进行rehash的触发会有用到的一个指标
	used int64
}

// hash表中的单个节点
type DictEntry struct {
	key * string // 假设string ，不然关联太多

	// 这个可以这么理解，val 是一个可以是其他类型（interface）的内容，也可以是int64 和uint64类型的
	// 内容，由于c语言可以进行这样操作，写go出来理解的话只能这样进行表述
	val *struct{ // redis 不支持泛型，所以只能一个个的标明
		v interface{}
		vui uint64
		vi int64
	}

	next *DictEntry
}

// 字段的操作函数，使用interface进行解耦
type DictType interface {
	// 进行计算hash函数的
	HashFunction(key string) int
	// 进行key的拷贝
	keyDup(*privdata )
	// 进行val 的拷贝
	ValDup(*privdata )
	// 进行key的对比
	KeyCompare()
	// 进行key的销毁
	KeyDestructor()
	// 进行值的销毁
	ValDestructor()
}