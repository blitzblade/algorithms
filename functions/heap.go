package functions

type HeapItem struct {
	parentIndex int
	leftIndex int
	rightIndex int
}

func GetHeapItem(arr []int, index int) HeapItem {
	return HeapItem{
		parentIndex: arr[index/2],
		leftIndex: arr[2*index],
		rightIndex: arr[2*index + 1],
	}
}