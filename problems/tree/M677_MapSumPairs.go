package tree

type msNode struct {
	sum   int
	nodes [26]*msNode
}

type MapSum struct {
	root *msNode
	mSum map[string]int
}

func NewMapSum() MapSum {
	return MapSum{
		root: &msNode{},
		mSum: make(map[string]int),
	}
}

func (ms *MapSum) Insert(key string, val int) {
	delta := val
	if ms.mSum[key] > 0 {
		delta -= ms.mSum[key]
	}
	ms.mSum[key] = val

	curr := ms.root
	for i := 0; i < len(key); i++ {
		idx := key[i] - 'a'
		if curr.nodes[idx] == nil {
			curr.nodes[idx] = &msNode{
				sum: val,
			}
		} else {
			curr.nodes[idx].sum += delta
		}
		curr = curr.nodes[idx]
	}

}

func (ms *MapSum) Sum(prefix string) int {
	curr := ms.root
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		curr = curr.nodes[idx]
		if curr == nil {
			return 0
		}
	}
	return curr.sum
}
