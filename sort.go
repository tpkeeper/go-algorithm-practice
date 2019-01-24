//排序相关算法
package main



//快排
func SortQuick(list []int) {
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

	SortQuick(list[0:mid])  //左
	SortQuick(list[mid+1:]) //右
}

//冒泡排序
func SortBubble(list []int) {
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


//堆排序
func SortHeap(list []int) {
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
