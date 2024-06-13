package lcr

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	t := &Trie3{
		children: [26]*Trie3{},
	}
	for _, v := range dictionary {
		t.Insert(v)
	}

	sArr := strings.Split(sentence, " ")
	for i := range sArr {
		sArr[i] = t.Search(sArr[i])
	}

	return strings.Join(sArr, " ")
}

type Trie3 struct {
	children [26]*Trie3
	end      bool
	value    string
}

func (this *Trie3) Insert(s string) {

	current := this

	for i, v := range s {
		if current.children[v-'a'] == nil {
			current.children[v-'a'] = &Trie3{
				children: [26]*Trie3{},
			}
		}
		if i == len(s)-1 {
			current.children[v-'a'].end = true
			current.children[v-'a'].value = s
		}

		current = current.children[v-'a']
	}

}

func (this *Trie3) Search(s string) string {
	current := this

	for _, v := range s {
		if current.end {
			return current.value
		}
		if current.children[v-'a'] == nil {
			return s
		}

		current = current.children[v-'a']

	}
	return s

}
