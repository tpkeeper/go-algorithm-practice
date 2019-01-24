package main

import "testing"

func TestReverseLinkList(t *testing.T) {
	node1:=&Node{"node1",nil}
	node2:=&Node{"node2",node1}
	node3:=&Node{"node3",node2}
	t.Log(node3.toString())
	LinkReverse(node3)
	t.Log(node1.toString())
}
