package test

import (
	"log"
	"testing"

	"gotest.tools/assert"
)

func TestInteger(t *testing.T) {
	var (
		a int
		b int8
		c int32
		u uint
	)

	a = 1
	b = int8(8)
	c = int32(32)
	u = uint(99)

	log.Printf("\n a:%d \n b:%d \n c:%d", a, b, c)

	assert.Equal(t, 1, a, "a not equal 1")
	assert.Equal(t, int8(8), b, "b not equal 8")
	assert.Equal(t, uint(99), u)
}
