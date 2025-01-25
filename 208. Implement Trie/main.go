package main

type Trie struct {
	nodes map[byte]*Trie
	word  bool
}

func Constructor() Trie {
	return Trie{
		nodes: make(map[byte]*Trie, 25),
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.word = true
		return
	}

	n, ok := this.nodes[word[0]]
	if ok {
		n.Insert(word[1:len(word)])
		return
	}

	x := Constructor()
	this.nodes[word[0]] = &x
	this.nodes[word[0]].Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.word
	}

	if n, ok := this.nodes[word[0]]; ok {
		return n.Search(word[1:])
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	if n, ok := this.nodes[prefix[0]]; ok {
		return n.StartsWith(prefix[1:])
	}

	return false
}

func main() {
	obj := Constructor()
	obj.Insert("apple")
	println(obj.Search("apple"))
	println(obj.StartsWith("apple"))
	println(obj.Search("app"))
	obj.Insert("app")
	println(obj.Search("app"))
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
