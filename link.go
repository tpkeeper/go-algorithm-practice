//链表相关算法
package main

type Node struct {
	value string
	next  *Node
}

func (node *Node) toString() (ret string){
	for node != nil {
		ret+=node.value+" -> "
		node = node.next
	}
	return
}
//反转链表
func LinkReverse(node *Node) {
	var pre *Node
	var now *Node
	var next *Node
	now = node

	for now != nil {
		// fmt.Println(now.value)
		next = now.next //备份下一个节点
		now.next = pre  //连接之前的节点

		pre = now //往后移动
		now = next
	}

}