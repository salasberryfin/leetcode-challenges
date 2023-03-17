package main

import (
	"testing"
)

//func TestCase1(t *testing.T) {
//	// construct new instance of Trie
//	trie := Constructor()
//
//	/**
//	insert 'apple'
//	search 'apple'
//	**/
//	// insert word 'apple'
//	trie.Insert("apple")
//
//	// check for inserted word 'apple'
//	existsApple := trie.Search("apple")
//	expectedApple := true
//	if existsApple != expectedApple {
//		t.Fatalf("trie.Search(\"apple\") = %v, expected %v\n", existsApple, expectedApple)
//	}
//
//	/**
//	search 'app'
//	startsWith 'app'
//	**/
//	// check for non-inserted word 'app'
//	existsApp := trie.Search("app")
//	expectedApp := false
//	if existsApp != expectedApp {
//		t.Fatalf("trie.Search(\"app\") = %v, expected %v\n", existsApp, expectedApp)
//	}
//
//	// check for words starting with 'app'
//	startsWithApp := trie.StartsWith("app")
//	expectedStartsWithApp := true
//	if startsWithApp != expectedStartsWithApp {
//		t.Fatalf("trie.StartsWith(\"app\") = %v, expected %v\n", startsWithApp, expectedStartsWithApp)
//	}
//
//	/**
//	startsWith 'rand'
//	**/
//	// check for words starting with 'rand'
//	startsWithRand := trie.StartsWith("rand")
//	expectedStartsWithRand := false
//	if startsWithRand != expectedStartsWithRand {
//		t.Fatalf("trie.StartsWith(\"rand\") = %v, expected %v\n", startsWithRand, expectedStartsWithRand)
//	}
//
//	/**
//	insert 'app'
//	search 'app'
//	search 'apps'
//	**/
//	// insert word 'app'
//	trie.Insert("app")
//
//	// check for inserted word 'app'
//	existsApp = trie.Search("app")
//	expectedApp = true
//	if existsApp != expectedApp {
//		t.Fatalf("trie.Search(\"app\") = %v, expected %v\n", existsApp, expectedApp)
//	}
//
//	// check for non-inserted word 'apps'
//	existApps := trie.Search("apps")
//	expectedApps := false
//	if existApps != expectedApps {
//		t.Fatalf("trie.Search(\"apps\") = %v, expected %v\n", existApps, expectedApps)
//	}
//
//}

func TestCase2(t *testing.T) {
	// construct new instance of Trie
	trie := Constructor()

	/**
	insert 'app'
	insert 'apple'
	insert 'beer'
	insert 'add'
	insert 'jam'
	insert 'rental'
	**/
	inserts := []string{
		"app",
		"apple",
		"beer",
		"add",
		"jam",
		"rental",
	}
	for _, i := range inserts {
		trie.Insert(i)
	}

	searches := map[string]bool{
		"apps":     false,
		"app":      true,
		"ad":       false,
		"applepie": false,
		"rest":     false,
		"jan":      false,
		"rent":     false,
		"beer":     true,
		"jam":      true,
	}

	for k, v := range searches {
		exist := trie.Search(k)
		if exist != v {
			t.Fatalf("trie.Search(\"%v\") = %v, expected %v\n", k, exist, v)
		}
	}
}
