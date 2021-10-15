package main

import "fmt"

// 具名函数
func Add(a, b int) int {
	return a + b
}

// 匿名函数
var add = func(a, b int) int {
	return a + b
}

// 多个返回值
func swap(a, b int) (int, int) {
	return b, a
}

// defer在return之后延迟执行匿名函数
func Inc() (v int) {
	defer func() {
		v++
	}()
	return 42
}

// 返回值命名
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

func main() {
	fmt.Println(add(1, 2))

	var a = []interface{}{123, "abc"}
	fmt.Println(a)
	j, k := swap(1, 2)
	fmt.Println(j, k)

	fmt.Println(Inc())
	fmt.Println(Find(map[int]int{1: 1, 2: 2}, 1))

	for i := 0; i < 3; i++ {
		defer func() {
			// 闭包对捕获的变量并不是传值方式，而是引用。
			println(i)
		}()

		// Output:
		// 3
		// 3
		// 3
	}

	for i := 0; i < 3; i++ {
		// 定义一个循环体内局部变量i
		i := i
		defer func() {
			println(i)
		}()
	}

	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) {
			println(i)
		}(i)
	}
}
