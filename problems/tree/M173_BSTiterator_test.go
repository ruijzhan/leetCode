package tree

import (
	"testing"
)

func TestIterator(t *testing.T) {
	itr := Constructor2(New([]int{7, 3, 15, null, null, 9, 20}))
	if itr.Next() != 3 {
		t.Fatal()
	}
	if itr.Next() != 7 {
		t.Fatal()
	}
	if !itr.HasNext() {
		t.Fatal()
	}
	if itr.Next() != 9 {
		t.Fatal()
	}
	if !itr.HasNext() {
		t.Fatal()
	}
	if itr.Next() != 15 {
		t.Fatal()
	}
	if !itr.HasNext() {
		t.Fatal()
	}
	if itr.Next() != 20 {
		t.Fatal()
	}
	if itr.HasNext() {
		t.Fatal()
	}
}
