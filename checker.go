package main

type Checker interface {
	Check(string) bool
}

type checker struct {
	root *trieNode
}

func newChecker(skipWords ...string) Checker {

	root := newTrieNode('/')
	sw := newSensitiveWords(skipWords...)

	// insert all sensitive words
	for _, word := range sw.words {
		root.insert(word)
	}

	return &checker{
		root: root,
	}
}

func (c *checker) Check(word string) bool {
	return c.root.find(word)
}
