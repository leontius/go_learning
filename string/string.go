package string

import (
	"log"
	"strings"
)

// Split hello world
func Split() {
	var hello = "hello,world!"
	// 切割
	var sArray = strings.Split(hello, ",")
	log.Printf("%v", sArray)
}

// EqualFold Example
// EqualFold 函数，计算 s 与 t 忽略字母大小写后是否相等。
func EqualFold() {
	log.Printf("String EqualFold：%t", strings.EqualFold("Go", "go"))
}

// Compare Example
// Compare 函数，用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。
// 不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
func Compare() {
	var (
		hello = "hello"
		h     = byte(107) // k
	)
	comp := strings.Compare(hello, string(h))
	if comp != 0 {
		log.Printf("result = %d ,var hello not compare h", comp)
	}
}

// Printf Example
func Printf() {
	var (
		a     = "字符串"
		b     = "\" 转义"
		hello = ""
	)

	log.Printf("%s,%s", a, b)

	hello = "hello, world!"

	h := hello[0]
	log.Printf("%+v", h)
	log.Printf("%s", string(h))
}

// Trim Example
func Trim() {
	var hello = "hello"
	hello = strings.Trim(hello, "h")
	log.Printf("%s", hello)
}

// Contains Example
// 子串 substr 在 s 中，返回 true
func Contains() {
	var hello = "hello"
	var h = "h"
	log.Printf("String Contains: %t", strings.Contains(hello, h))
}

// ContainsAny Example
// chars 中任何一个 Unicode 代码点在 s 中，返回 true
func ContainsAny() {
	var hello = "hello"
	var c = "h"
	log.Printf("String ContainsAny: %t", strings.ContainsAny(hello, c))
}

// ContainsRune Example
// Unicode 代码点 r 在 s 中，返回 true
func ContainsRune() {
	var hello = "hello"
	var c = rune(0x1fffffff)
	log.Printf("String ContainsRune: %t %s %d", strings.ContainsRune(hello, c), hello, c)
}
