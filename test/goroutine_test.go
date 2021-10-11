package test

import (
	"testing"

	"github.com/leontius/fetch"
)

func TestFetch(t *testing.T) {
	fetch.New("https://www.sina.com", "https://bing.com", "https://baidu.com")
}
