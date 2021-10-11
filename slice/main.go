package main

import "fmt"

var (
	// nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
	a []int
	// 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	b = []int{}
	// 有3个元素的切片, len和cap都为3
	c = [...]int{1, 2, 3}
	// 有2个元素的切片, len为2, cap为3
	d = c[:2]
	// 有2个元素的切片, len为2, cap为3
	e = c[0:2:cap(c)]
	// 有0个元素的切片, len为0, cap为3
	f = c[:0]
	// 有3个元素的切片, len和cap都为3
	g = make([]int, 3)
	// 有3个元素的切片，len为2，cap为3
	h = make([]int, 2, 3)
	// 有0个元素的切片，len为0，cap为3
	i = make([]int, 0, 3)
)

func main() {
	fmt.Println(a, b)
	fmt.Println(c, cap(c))
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	// 内置函数len可以用于计算数组的长度，cap函数可以用于计算数组的容量。
	fmt.Println(g, len(g), cap(g))
	// cap函数只计算左指针到原array最后的值的个数
	h[0] = 1
	fmt.Println(h, len(h), cap(h))
	fmt.Println(i)

	// 遍历切片的方式和遍历数组的方式类似
	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}

	// 追加1个元素
	a = append(a, 1)
	fmt.Println(a)
	// 追加多个元素
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)
	// 追加一个切片，切片需要解包
	a = append(a, []int{1, 2, 3, 4, 5}...)
	fmt.Println(a)
	// 在开头追加一个元素
	a = append([]int{1}, a...)
	// 在开头追加一个切片
	a = append([]int{-1, -2, -3}, a...)
	fmt.Printf("%v 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因此性能较尾部追加来的差。\n", a)

	// append支持链式调用

	i := 1
	x := 999
	// 从i个位置插入x
	a = append(a[:i], append([]int{x}, a[i:]...)...)
	fmt.Println(a)
	// 从i个位置插入切片
	a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...)
	fmt.Println(a)
}
