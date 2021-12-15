package designaddandsearchwordsdatastructure

const leafNode = byte(255)

type WordDictionary struct {
	dict map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{dict: make(map[byte]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	addWord(this, word)
}

func addWord(wDict *WordDictionary, word string) {
	if len(word) == 0 {
		wDict.dict[leafNode] = nil
		return
	}
	if _, ok := wDict.dict[word[0]]; !ok {
		wDict.dict[word[0]] = &WordDictionary{dict: make(map[byte]*WordDictionary)}
	}
	addWord(wDict.dict[word[0]], word[1:])
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	return search(this, word)
}

func search(wDict *WordDictionary, word string) bool {
	if len(word) == 0 {
		_, existLeafNode := wDict.dict[leafNode]
		return existLeafNode
	}
	ret := false
	if word[0] == (".")[0] {
		for k, v := range wDict.dict {
			if k == leafNode {
				continue
			}
			ret = ret || search(v, word[1:])
		}
	} else if v, ok := wDict.dict[word[0]]; ok {
		ret = search(v, word[1:])
	}
	return ret
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
