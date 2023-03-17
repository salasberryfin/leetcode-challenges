/*
*
This is a very interesting problem

A Trie can be implemented as:
  - Hashmap {byte: *Trie}
  - Array [26]*Trie - for each letter in the alphabet

Array is faster and takes less memory: especially effective if strings are short
Hash map does not occupy as much unused memory and is effective for large strings or
some characters occur very rarely
- With an array, we're reserving space for 26 elements even if we never find a specific character
- With a hash map, each element uses more space, but we only take the space of the characters we insert

While implementing the insert:
  - Be careful with initializing each Trie node -> line 31

*
*/
package main

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	Children map[byte]*Trie
	Value    string
	Last     bool
}

func Constructor() Trie {
	return Trie{
		Children: make(map[byte]*Trie),
		Last:     false,
	}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, i := range word {
		if current.Children[byte(i)] == nil {
			current.Children[byte(i)] = &Trie{
				Children: make(map[byte]*Trie),
				Last:     false,
			}
		}
		current = current.Children[byte(i)]
	}
	current.Value = word
	current.Last = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for _, i := range word {
		if current.Children[byte(i)] == nil {
			return false
		}
		current = current.Children[byte(i)]
	}

	return current.Last
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, i := range prefix {
		if current.Children[byte(i)] == nil {
			return false
		}
		current = current.Children[byte(i)]
	}

	return true
}
