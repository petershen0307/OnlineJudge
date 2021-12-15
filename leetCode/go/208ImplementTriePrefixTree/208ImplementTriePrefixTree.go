package implementtrieprefixtree

const leafNode = byte(255)

type Trie struct {
	dict map[byte]*Trie
}

func Constructor() Trie {
	return Trie{dict: map[byte]*Trie{}}
}

func (this *Trie) Insert(word string) {
	insert(this, word)
}
func insert(wDict *Trie, word string) {
	if len(word) == 0 {
		wDict.dict[leafNode] = nil
		return
	}
	if _, ok := wDict.dict[word[0]]; !ok {
		wDict.dict[word[0]] = &Trie{dict: make(map[byte]*Trie)}
	}
	insert(wDict.dict[word[0]], word[1:])
}
func (this *Trie) Search(word string) bool {
	return search(this, word)
}

func search(wDict *Trie, word string) bool {
	if len(word) == 0 {
		_, existLeafNode := wDict.dict[leafNode]
		return existLeafNode
	}
	ret := false
	if v, ok := wDict.dict[word[0]]; ok {
		ret = search(v, word[1:])
	}
	return ret
}

func (this *Trie) StartsWith(prefix string) bool {
	return startsWith(this, prefix)
}

func startsWith(wDict *Trie, word string) bool {
	if len(word) == 0 {
		return true
	}
	ret := false
	if v, ok := wDict.dict[word[0]]; ok {
		ret = startsWith(v, word[1:])
	}
	return ret
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
