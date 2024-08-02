// ARBOL BUSQUEDA DE PREFIJOS
package main

import "fmt"

type Node struct {
	children map[rune]*Node
	wordEnd  bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{children: make(map[rune]*Node)},
	}
}

func (t *Trie) Insert(word string) {
	actualNode := t.root
	for _, ch := range word {
		if _, exists := actualNode.children[ch]; !exists {
			actualNode.children[ch] = &Node{children: make(map[rune]*Node)}
		}
		actualNode = actualNode.children[ch]
	}
	actualNode.wordEnd = true
}

func (t *Trie) StartsWith(word string) bool {
	actualNode := t.root
	for _, ch := range word {
		if _, exists := actualNode.children[ch]; !exists {
			return false
		}
		actualNode = actualNode.children[ch]
	}
	return true
}

func (t *Trie) Search(word string) bool {
	actualNode := t.root
	for _, ch := range word {
		if _, exists := actualNode.children[ch]; !exists {
			return false
		}
		actualNode = actualNode.children[ch]
	}
	return actualNode.wordEnd
}

func main() {
	trie := NewTrie()
	words := []string{"hola", "mundo", "holaquetal", "holanda"}

	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("hola"))
	fmt.Println(trie.StartsWith("m"))
}
