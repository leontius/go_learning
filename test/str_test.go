package test

import (
	"testing"

	"github.com/leontius/string"
)

func TestString(t *testing.T) {
	// 打印
	string.Printf()

	// 比较
	string.Compare()

	// 指定去重
	string.Compare()

	// 字符串分割
	string.Split()

	// 字符串比较
	string.EqualFold()

	// 包含
	string.Contains()
	string.ContainsAny()
	string.ContainsRune()
}
