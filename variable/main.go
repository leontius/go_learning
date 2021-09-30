package main

import (
	"fmt"
	"io"
	"os"
)

func zeroValue() {
	var s string
	// output: ""
	fmt.Print(s)
}

func initValue() {
	var i, j, k int
	fmt.Printf("%d %d %d \n", i, j, k)

	var b, d, s = true, 2.3, "four"
	fmt.Printf("%t %f %s \n", b, d, s)

	// os.Open returns a file and an error
	f, err := os.Open("/tmp")
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v\n", f)
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			// log etc
			fmt.Print(err)
		}
	}(f)
}

func pointer() {
	x := 1
	// p, of type *int, points to x
	p := &x
	fmt.Println(*p)
	// equivalent to x = 2
	*p = 2
	fmt.Printf("x:%v p:%v = %d\n", &x, &p, *p)

	var j, k int
	// "true false false"
	fmt.Println(&j == &j, &j == &k, &j == nil)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	// 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	*p++
	return *p
}

// new函数类似是一种语法糖
// newInt1 和 newInt2 行为相同
func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	// zero value
	zeroValue()

	// init
	initValue()

	// pointer
	pointer()

	// return pointer
	var p = f()
	fmt.Println(*p)
	// "false"
	fmt.Println(f() == f())

	v := 1
	// side effect: v is now 2
	incr(&v)
	// "3" (and v is 3)
	fmt.Println(incr(&v))

	newInt1()
	newInt2()

	// 每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的
	// 特殊情况：如果两个类型都是空的，也就是说类型的大小是0，有可能返回相同地址
	j, k := new(int), new(int)
	fmt.Println(j == k)
}
