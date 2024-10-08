package main

type trieNode struct {
	isEnd    bool
	char     rune
	children map[rune]*trieNode
}

func newTrieNode(char rune) *trieNode {
	return &trieNode{
		isEnd:    false,
		char:     char,
		children: make(map[rune]*trieNode),
	}
}

func (t *trieNode) insert(word string) {

	node := t
	for _, c := range word {

		if v, ok := node.children[c]; ok {
			node = v
			continue
		} else {
			tn := newTrieNode(c)
			node.children[c] = tn
			node = tn
		}

	}

	node.isEnd = true
}

func (t *trieNode) find(word string) bool {

	node := t
	for _, c := range word {

		if v, ok := node.children[c]; !ok {
			return false
		} else {
			node = v
		}
	}

	return node.isEnd
}
