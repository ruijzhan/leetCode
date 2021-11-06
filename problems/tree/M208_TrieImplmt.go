package tree

type Trie struct {
	char    byte
	nodes   [26]*Trie
	wordEnd bool
}

func NewTrie() Trie {
	return Trie{char: '/'}
}

func (t *Trie) Insert(word string) {
	curr := t
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curr.nodes[index] == nil {
			curr.nodes[index] = &Trie{char: word[i]}
		}
		curr = curr.nodes[index]
	}
	curr.wordEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curr.nodes[index] == nil {
			return false
		}
		curr = curr.nodes[index]
	}
	return curr.wordEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if curr.nodes[index] == nil {
			return false
		}
		curr = curr.nodes[index]
	}
	return true
}
