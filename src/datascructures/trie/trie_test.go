package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Add(t *testing.T) {
	trie := NewTrie()
	err := trie.Add("bruno")
	assert.Nil(t, err)
}

func Test_Search(t *testing.T) {
	trie := NewTrie()
	err := trie.Add("fatima")
	assert.Nil(t, err)

	err = trie.Add("bruno")
	assert.Nil(t, err)

	err = trie.Add("bruna")
	assert.Nil(t, err)

	err = trie.Add("breno")
	assert.Nil(t, err)

	err = trie.Add("bru")
	assert.Nil(t, err)

	founds := trie.Search("fatima")
	assert.Equal(t, []string{"fatima"}, founds)

	founds = trie.Search("bruno")
	assert.Equal(t, []string{"bruno"}, founds)

	founds = trie.Search("br")
	assert.Equal(t, []string{"breno", "bru", "bruna", "bruno"}, founds)
}

func Test_Exists(t *testing.T) {
	trie := NewTrie()
	err := trie.Add("fatima")
	assert.Nil(t, err)

	err = trie.Add("bruno")
	assert.Nil(t, err)

	err = trie.Add("bruna")
	assert.Nil(t, err)

	err = trie.Add("breno")
	assert.Nil(t, err)

	err = trie.Add("bru")
	assert.Nil(t, err)

	exists := trie.Exists("fatima")
	assert.Equal(t, true, exists)

	exists = trie.Exists("fatim")
	assert.Equal(t, false, exists)

	exists = trie.Exists("bruno")
	assert.Equal(t, true, exists)

	exists = trie.Exists("bruna")
	assert.Equal(t, true, exists)

	exists = trie.Exists("breno")
	assert.Equal(t, true, exists)

	exists = trie.Exists("bru")
	assert.Equal(t, true, exists)

	exists = trie.Exists("br")
	assert.Equal(t, false, exists)
}

func Test_Delete(t *testing.T) {
	trie := NewTrie()

	err := trie.Add("fatima")
	assert.Nil(t, err)

	exists := trie.Exists("fatima")
	assert.Equal(t, true, exists)

	trie.Remove("fatima")

	exists = trie.Exists("fatima")
	assert.Equal(t, false, exists)

	err = trie.Add("fatima")
	assert.Nil(t, err)

	exists = trie.Exists("fatima")
	assert.Equal(t, true, exists)
}
