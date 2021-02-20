## 数组
数组是一种线性数据结构，它用一组连续的内存空间，来存储一组具有相同类型的数据。
在go语言中，我们会从类型、大小两个方面去描述数组。类型相同，大小不同的数组在go语言看来是完全不同的，甚至不能用 `==` 比较。

```go
func NewArray(elem *Type, bound int64) *Type {
	if bound < 0 {
		Fatalf("NewArray: invalid bound %v", bound)
	}
	t := New(TARRAY)
	t.Extra = &Array{Elem: elem, Bound: bound}
	t.SetNotInHeap(elem.NotInHeap())
	return t
}
```
go语言编译期间生成数组的代码位于
[cmd/compile/internal/types.NewArray](https://draveness.me/golang/tree/cmd/compile/internal/types.NewArray)
。数组类型包含2个字段，`bound`确定大小，`Elem`确定元素类型。

更多内容可阅读：[数组编译原理](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array/)

## 数组的查找、插入、删除复杂度
根据数组下标访问或赋值的时间复杂度都是O(1)。

无序数组：
查找复杂度：无序数组遍历查找，复杂度为O(n)  
插入复杂度：在空间足够的情况下，插入复杂度为O(1)，空间不足情况下为O(n)
删除复杂度：查找复杂度为O(n)，赋值复杂度为O(1)，因此删除的复杂度为O(n)

有序数组：
查找复杂度：O(log2(n))，基于二分查找算法，后面章节会详细说明。
插入复杂度：为保证数组有序性，插入位置之后的元素将后移一位，复杂度为O(n) + O(log2(n))，即：O(n)
删除复杂度：查找复杂度为O(log2(n))，删除位置后的所有元素前移一位，复杂度为O(n) + O(log2(n))，即：O(n)

## 动态数组（Slice）
在go语言中，数组创建之初就已经固定了大小，无法扩容，从而设计了依托数组的动态数组：切片（Slice）。

# NEXT
[动态数组（Slice）](../b_动态数组Slice)

