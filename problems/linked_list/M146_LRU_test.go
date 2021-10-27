package linkedlist

import (
	"testing"
)

func TestLRU(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	if lru.Get(1) != 1 {
		t.Fatal()
	}

	lru.Put(3, 3)
	if lru.Get(2) != -1 {
		t.Fatal()
	}
	lru.Put(4, 4)
	if lru.Get(1) != -1 {
		t.Fatal()
	}
	if lru.Get(3) != 3 {
		t.Fatal()
	}
	if lru.Get(4) != 4 {
		t.Fatal()
	}
}

func TestLRU2(t *testing.T) {
	lru := Constructor(2)
	if lru.Get(2) != -1 {
		t.Fatal()
	}
	lru.Put(2, 6)
	if lru.Get(1) != -1 {
		t.Fatal()
	}
	lru.Put(1, 5)
	lru.Put(1, 2)
	if lru.Get(1) != 2 {
		t.Fatal()
	}
	if lru.Get(2) != 6 {
		t.Fatal()
	}
}
