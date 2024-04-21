package topics


/*
 前缀树 模板
*/
type Trie struct {
	childrenMap map[byte]*Trie
	children    []*Trie
	end         bool
	value       int
}

func Constructor() Trie {
	dummy := Trie{
		childrenMap: make(map[byte]*Trie, 0),
		children:    make([]*Trie, 0),
		end:         false,
		value:       -1,
	}

	return dummy
}

func (this *Trie) Insert(word string) {
	current := this
	for i := 0; i < len(word); i++ {
		if _, ok := current.childrenMap[word[i]]; !ok {
			trie := &Trie{
				childrenMap: make(map[byte]*Trie, 0),
				children:    make([]*Trie, 0),
				end:         false,
				value:       int(word[i]),
			}

			current.childrenMap[word[i]] = trie
			current.children = append(current.children, trie)
		}

		if i == len(word)-1 {
			current.childrenMap[word[i]].end = true
		}
		current = current.childrenMap[word[i]]
	}

}

func (this *Trie) Search(word string) bool {

	current := this

	for i := 0; i < len(word); i++ {
		if _, ok := current.childrenMap[word[i]]; !ok {
			return false
		}
		trie := current.childrenMap[word[i]]

		if i == len(word)-1 && !trie.end {
			return false
		}
		current = current.childrenMap[word[i]]
	}
	return true

}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for i := 0; i < len(prefix); i++ {
		if _, ok := current.childrenMap[prefix[i]]; !ok {
			return false
		}
		current = current.childrenMap[prefix[i]]
	}
	return true
}

