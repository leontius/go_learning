package q4

import (
	"fmt"

	"github.com/leontius/q4/lib"
	// "github.com/leontius/q4/lib/internal"
	// 无法引用internal包，internal包只能被直接父级和子级包引用
)

func DemoQ4() {
	fmt.Println("q4")
	lib.Demo7()
}
