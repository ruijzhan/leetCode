package tree

import (
	"testing"
)

func TestNewMapSum(t *testing.T) {
	ms := NewMapSum()
	ms.Insert("apple", 3)
	if ms.Sum("ap") != 3 {
		t.Fatal()
	}
	ms.Insert("app", 2)
	if ms.Sum("ap") != 5 {
		t.Fatal()
	}

	ms = NewMapSum()
	ms.Insert("apple", 3)
	if ms.Sum("ap") != 3 {
		t.Fatal()
	}
	ms.Insert("app", 2)
	ms.Insert("apple", 2)
	if ms.Sum("ap") != 4 {
		t.Fatal()
	}
}
