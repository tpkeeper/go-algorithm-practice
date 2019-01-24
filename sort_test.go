package main

import "testing"

func TestSortBubble(t *testing.T) {
	list := []int{4, 67, 8, 1, 45, 49, 2, 43, 88,49}
	t.Log(list)
	SortBubble(list)
	t.Log(list)
}

func TestSortQuick(t *testing.T) {
	list := []int{4, 67, 8, 1, 45, 49, 2, 43, 88,49}
	t.Log(list)
	SortQuick(list)
	t.Log(list)
}

func TestSortHeap(t *testing.T) {
	list := []int{4, 67, 8, 1, 45, 49, 2, 43, 88,49}
	t.Log(list)
	SortHeap(list)
	t.Log(list)
}