package test

import (
	"log"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
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

	h = byte(01)
	if !strings.Contains(hello, string(h)) {
		log.Print("var hello not contains h")
	}

	comp := strings.Compare(hello, string(h))
	if comp != 0 {
		log.Printf("result = %d ,var hello not compare h", comp)
	}

	hello = strings.Trim(hello, "h")
	log.Printf("%s", hello)

	var sArray = strings.Split(hello, ",")
	log.Printf("%v", sArray[0])

	// print char
	for i := 0; i < len(hello); i++ {
		log.Printf("%v", string(hello[i]))
	}
}
