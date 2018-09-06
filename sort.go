package main

import (
	"fmt"
)

func main() {
	list := []int{4, 67, 8, 1, 45, 49, 2, 43, 88}
	fmt.Println("冒泡排序")
	listBubble := make([]int, len(list))
	copy(listBubble, list)
	fmt.Println(listBubble)
	sortBubble(listBubble)
	fmt.Println(listBubble)

	fmt.Println("快速排序")
	listFast := make([]int, len(list))
	copy(listFast, list)
	fmt.Println(listFast)
	sortFast(listFast)
	fmt.Println(listFast)

	fmt.Println("堆排序")
	listHeap := make([]int, len(list))
	copy(listHeap, list)
	fmt.Println(listHeap)
	// adjustHeap(listHeap, len(listHeap)-1)
	sortHeap(listHeap)
	fmt.Println(listHeap)

	fmt.Println("链表反序")
	n1 := Node{
		v: "n1",
	}
	n2 := Node{
		v: "n2",
		p: &n1,
	}
	n3 := Node{
		v: "n3",
		p: &n2,
	}

	n3.print()
	reverseLinkList(&n3)
	n1.print()

}

type Node struct {
	v string
	p *Node
}

func (node *Node) print() {
	fmt.Println()
	for node != nil {
		fmt.Print(node.v)
		node = node.p
	}
}

func reverseLinkList(node *Node) {
	var pre *Node
	var now *Node
	var next *Node

	now = node

	for now != nil {
		// fmt.Println(now.v)
		next = now.p //备份下一个节点
		now.p = pre  //连接之前的节点

		pre = now //往后移动
		now = next
	}

}

func sortFast(list []int) {
	if len(list) == 0 {
		return
	}
	mid := 0         //tag索引
	tag := list[mid] //tag
	i, j := 1, len(list)-1
	for mid < j { //最后一次比较为mid=j-1
		if list[i] > tag { //大的放到最右
			list[i], list[j] = list[j], list[i] //交换
			j--                                 //更新待比较区域右边
		} else { //小的放到自己左边
			list[i], list[mid] = list[mid], list[i]
			mid++ //更新tag索引
			i++   //更新待比较区域左边
		}
		// fmt.Println(list)
	}

	sortFast(list[0:mid])  //左
	sortFast(list[mid+1:]) //右
}

func sortBubble(list []int) {
	len := len(list)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}

	}
}

func sortHeap(list []int) {
	for i := len(list) - 1; i >= 2; i-- { //每次调整完之后，都将头尾互调
		adjustHeap(list, i)
		list[0], list[i] = list[i], list[0]
	}
}

func adjustHeap(list []int, tail int) {
	len := tail + 1
	for i := len/2 - 1; i >= 0; i-- { //循环调整每一个非叶子节点，最后形成大堆
		if (2*i+2 <= tail) && (list[2*i+1] < list[2*i+2]) {
			if list[i] < list[2*i+2] {
				list[i], list[2*i+2] = list[2*i+2], list[i]
			}

		} else {
			if list[i] < list[2*i+1] {
				list[i], list[2*i+1] = list[2*i+1], list[i]
			}
		}
	}
}
