package main

type Checker interface {
	Check(string) bool
}

type checker struct {
	root *trieNode
}

func newChecker(skipWords ...string) (Checker, error) {

	root := newTrieNode('/')
	sw, err := newSensitiveWords(skipWords...)
	if err != nil {
		return nil, err
	}

	// insert all sensitive words
	for _, word := range sw.words {
		root.insert(word)
	}

	return &checker{
		root: root,
	}, nil
}

func (c *checker) Check(word string) bool {
	return c.root.find(word)
}
