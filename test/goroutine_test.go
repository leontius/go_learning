package test

import (
	"testing"

	"github.com/leontius/fetch"
)

func TestFetch(t *testing.T) {
	urls := []string{"https://www.sina.com", "https://bing.com", "https://baidu.com"}
	fetch.New(urls)
}
