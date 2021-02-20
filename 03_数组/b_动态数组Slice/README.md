## slice
上一节介绍的数组，在go语言中没那么常用，更常用的数据结构是切片，即动态数组，其长度并不固定，我们可以向切片中追加元素，
它会在容量不足时自动扩容。


## 结构

```go
func newslice(elem *type) *type {
	if t := elem.cache.slice; t != nil {
		if t.elem() != elem {
			fatalf("elem mismatch")
		}
		return t
	}

	t := new(tslice)
	t.extra = slice{elem: elem}
	elem.cache.slice = t
	return t
}
```

```go
// slice contains type fields specific to slice types.
type slice struct {
	elem *type // element type
}
```
[cmd/compile/internal/types.newslice](https://draveness.me/golang/tree/cmd/compile/internal/types.newslice)

结合上述代码片段我们可以发现，在编译期间slice只确定了类型，存储于`extra`中，而没有确定长度和容量，因为slice的长度和容量是可变的。

```go
// 运行时，/usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/slice.go
type slice struct {
	array unsafe.pointer
	len   int
	cap   int
}
```
slice的运行时结构体如下：  
1. array是一个指向数组的指针，表明了slice是依托于数组实现。
2. len表示slice的长度。
3. cap表示slice的容量，同时也是底层数组 array 的长度。

## 切片底层实现
网络上介绍切片的文章大把的都在，这里引用go语言中文网的文章，请大家移步观看。  
[Go语言切片详解](https://studygolang.com/articles/28914)

## append详解
go语言中可以对切片进行 `append` 操作，根据容量的大小将自动进行扩容。

1. 容量足够的情况下，进行append(slice, 1)操作，会将数据追加到原数组上对应的位置，不会创建新的数组。
2. 容量不够的情况下，进行append(slice, 1)操作，将创建一个新的底层数组，长度为原底层数组的2倍，同时slice中的array
指针指向新的数组。
3. 容量不够的情况下，进行append(slice, []int{1, 2, 3}...)操作，如果append的之后的元素数量大于之前容量的2倍，
那么新数组的长度为 原容量*2+（元素总数量-原容量*2+1）/ 2。  
例如原slice容量为3，长度为3，进行append(slice, 2, 2, 2, 2)操作后，slice的容量为8，而不是12。  
[参考示例](../../code/03/slice/append.go)

## Slice的缺点
我们知道slice是引用类型，array是值类型，按理来说引用传递的效率优于值传递。一定是这样吗？  
答案是否定的，我们做一个基准测试来说明：[slice vs array](../../code/03/sliceVsArray/compare_test.go)

从基准测试的结果可以看到：  
创建数组的循环次数480286，平均执行时间2465 ns/op，堆上分配内存数量为0，内存总量为0。   
创建Slice的循环次数355299，平均执行时间3472 ns/op，堆上分配内存数量为1，内存总量为8192B。

如此看来，并非所有情况都要使用slice，之所以产生这种情况是因为发生了`内存逃逸`。

# NEXT
[内存逃逸](../c_内存逃逸)