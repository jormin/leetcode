package problem

import (
	"fmt"
	"testing"
)

// TestLRUCache 测试LRU缓存
func TestLRUCache(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(2, 3)
	got := l.Get(1)
	if got != 1 {
		t.Errorf("get key error, got%v, want %v", got, 1)
	}
	l.Put(3, 3)
	l.Get(2)
	got = l.Get(2)
	if got != -1 {
		t.Errorf("get key error, got%v, want %v", got, -1)
	}
	l.Put(4, 4)
	l.Get(1)
	got = l.Get(1)
	if got != -1 {
		t.Errorf("get key error, got%v, want %v", got, -1)
	}
	l.Get(3)
	got = l.Get(3)
	if got != 3 {
		t.Errorf("get key error, got%v, want %v", got, 3)
	}
	l.Get(4)
	got = l.Get(4)
	if got != 4 {
		t.Errorf("get key error, got%v, want %v", got, 4)
	}
	fmt.Println(l.String())
}
