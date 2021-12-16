package memo

import (
	"testing"

	"github.com/sanopy/gobook/ch09/ex03/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestConcurrentCancel(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.ConcurrentCancel(t, m)
}
