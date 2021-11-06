package tree

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	if trie.Search("apple") != true {
		t.Fatal()
	}
	if trie.Search("app") == true {
		t.Fatal()
	}
	if trie.StartsWith("app") != true {
		t.Fatal()
	}
	trie.Insert("app")
	if trie.Search("app") != true {
		t.Fatal()
	}
}
