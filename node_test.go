package main

import (
	"testing"
)

var nodeRoot, nodeRoot2 *trieNode

func init() {
	nodeRoot = newTrieNode('/')

	nodeRoot2 = newTrieNode('/')
	nodeRoot2.insert("sensitive")
	nodeRoot2.insert("blood")
}

func Test_trieNode_insert(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		tr   *trieNode
		args args
	}{
		{
			"test1",
			nodeRoot,
			args{
				"hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.insert(tt.args.word)
		})
	}
}

func Test_trieNode_find(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		tr   *trieNode
		args args
		want bool
	}{
		{
			"word-sensitive",
			nodeRoot2,
			args{
				"sensitive",
			},
			true,
		},
		{
			"word-insensitive",
			nodeRoot2,
			args{
				"insensitive",
			},
			false,
		},
		{
			"word-sensitive-not",
			nodeRoot2,
			args{
				"sensitive-not",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.find(tt.args.word); got != tt.want {
				t.Errorf("trieNode.find() = %v, want %v", got, tt.want)
			}
		})
	}
}
