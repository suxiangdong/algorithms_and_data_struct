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


