package test

import (
	"os"
	"testing"

	"github.com/leontius/fetch"
)

func TestFetch(t *testing.T) {
	os.Args[1] = "https://www.sina.com"
	os.Args[2] = "https://bing.com"
	fetch.New()
}
