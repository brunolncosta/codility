package datascructures

import (
	"errors"
)

// Trie is a tree of words. It allows store and search word in O(M) where M is the size of the word.
// It also allows search for words using by a prefix witch is a O(N) where N is the size of the letter in the trie.
type Trie interface {
	Add(word string) error
	Remove(word string)
	Search(word string) []string
	Exists(word string) bool
}

type trieImpl struct {
	root Node
}

func NewTrie() Trie {
	return &trieImpl{
		root: newNode(0),
	}
}

func (trie trieImpl) Add(word string) error {
	if word == "" {
		return errors.New("invalid word")
	}

	letters := []rune(word)
	curr := &trie.root
	for _, letter := range letters {
		if child := curr.childNodes[letter]; child != nil {
			curr = child
			continue
		}

		node := newNode(letter)
		curr.childNodes[letter] = &node
		curr = curr.childNodes[letter]
	}

	curr.isWord = true
	return nil
}

func (trie trieImpl) Remove(word string) {
	letters := []rune(word)
	if node, exists := trie.getWordNode(letters); exists {
		node.isWord = false
	}
}

func (trie trieImpl) Search(word string) []string {
	founds := make([]string, 0)
	letters := []rune(word)
	nodeWord, exists := trie.getWordNode(letters)
	if exists {
		founds = append(founds, word)
		return founds
	}
	getAllWordRecursive(letters[:len(letters)-1], nodeWord, &founds)
	return founds
}

func (trie trieImpl) getWordNode(letters []rune) (*Node, bool) {
	curr := &trie.root
	for _, letter := range letters {
		if child := curr.childNodes[letter]; child != nil {
			curr = child
		}
	}
	return curr, curr.isWord
}

func getAllWordRecursive(base []rune, curr *Node, founds *[]string) {
	base = append(base, curr.letter)
	if curr.isWord {
		*founds = append(*founds, string(base))
	}

	if len(curr.childNodes) > 0 {
		for _, child := range curr.childNodes {
			if child != nil {
				getAllWordRecursive(base, child, founds)
			}
		}
	}
}

func (trie trieImpl) Exists(word string) bool {
	letters := []rune(word)
	_, exists := trie.getWordNode(letters)
	return exists
}

type Node struct {
	letter     rune
	childNodes []*Node // could be a map[rune]*Node as well, but the search result won't be sorted
	isWord     bool
}

func newNode(letter rune) Node {
	return Node{
		letter:     letter,
		childNodes: make([]*Node, 1000),
		isWord:     false,
	}
}
