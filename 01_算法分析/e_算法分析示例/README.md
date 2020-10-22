## 常量时间操作 O(1)
go语言中内置函数：len()
```go
package types
// A Type represents a Go type.
type Type struct {
    // ...
}
// Array contains Type fields specific to array types.
type Array struct {
	Elem  *Type // element type
	Bound int64 // number of elements; <0 if unknown yet
}
```
上面是go中，数组（[代码](https://github.com/golang/go/blob/master/src/cmd/compile/internal/types/type.go#L340 "查看地址")）
结构体的定义。我们知道数组是定长的，当获取数组的长度时，可以直接获取Array类型变量的Bound属性，而不需要遍历数组元素确定个数，所以我们时说复杂度
是O(1)。  
当查看len()函数时会发现，builtin.go文件中，只有对len()函数的定义，而没有具体实现。内置函数的解析是在中间代码生成阶段处理的。
[参考文章](https://mp.weixin.qq.com/s/iO5qjcCql-MPJiatUtdiHQ "")

## 线性时间操作 O(n)
[查找序列内最大值](../../code/01/max.go)  
我们可以看到，在循环体开始之前，原子操作数量最多为常数2，循环体内的原子操作执行数量与输入的切片长度有关，就是长度n，所以我们说
找最大值算法的时间复杂度为O(n)

## 对数时间操作 O(log(n))
在查找最大值算法的循环中，有一个max赋值操作。如果输入序列内的元素唯一，且升序排列，那么要进行n-1次赋值。如果是随机序列，那么第i个元素比前i个
元素更大的概率为1/i，所以max赋值操作次数预期是 1 + 1/2 + 1/3 + ... + 1/i，
这是著名的[调和数](https://baike.baidu.com/item/%E8%B0%83%E5%92%8C%E7%BA%A7%E6%95%B0/8019971?fr=aladdin "调和数")。
简单来说就是当n很大时，1 + 1/2 + 1/3 + ... + 1/i 有一个渐进表达式 log(n) + c ，c是常数。（具体原理不清楚，有兴趣可自行研究）  

对于赋值操作来说，时间复杂度并不是O(log(n))，算法研究的是最坏情况输入，明显是O(n)，比如序列长度n=2^100，并且序列升序。

时间复杂度为O(log(n))的函数:  
```go
package c
import "fmt"
func count (n int) {
    cnt := 0
    for i:= 1;i < n; i *= 2 {
        cnt++
    }
    fmt.Println(cnt)
}
```
在这个函数中，cnt就代表了原子操作执行数量，i = i * 2^cnt，那么最大执行次数就是当i=n时，所以就是求n = n * 2^cnt的解，可以得到
cnt = log(n)，即时间复杂度为O(log(n))。  
从这个函数我们可以想到二分查找算法，中位索引的取值为长度(n-1)/2, (n-1)/2/2 ... 1, 0, -1，与上边这个简单函数有些类似，
时间复杂度同样为O(log(n))，大家可以自己写一下二分查找验证一下。

## 二次时间操作 O(n^2)
前缀平均值问题：给定一个长度为n的序列S，计算出A序列，j < n时，A序列满足A[j] = (S[0] + s[1] + s[2] + ... + S[j]) / (j + 1)  
[前缀平均值](../../code/01/prefix_average.go)  
代码中给出了2种解法，对第一种解法来说，外层循环次数为n，内层循环次数为 1 + 2 + .. + n - 1。  
在介绍数学函数的时候，有一个数学定理：1+2+3+ ... + n = n(n+1)/2，那么这种算法的原子操作执行次数就是 n * ((n-1)*n/2)，
所以时间复杂度就是O(n^2)。  
第二种解法中，将保存每次计算得到的前j个值的和，不再重新计算前j个值的和，因此去掉了内层循环，时间复杂度为O(n)

