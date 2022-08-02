package functions

type HeapItem struct {
	parentIndex int
	leftIndex int
	rightIndex int
}

func GetHeapItem(arr []int, index int) HeapItem {

	return HeapItem{
		parentIndex: index/2,
		leftIndex: 2*index,
		rightIndex: 2*index + 1,
	}
}

func MaxHeapify(arr []int, i int) {
	l := 2*i
	r := 2*i + 1
	n := len(arr) - 1

	var largest int;

	if l <= n && arr[l] > arr[i] {
		largest = l
	}else{
		largest = i
	}

	if r <= n && arr[r] > arr[largest]{
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		MaxHeapify(arr, largest)
	}
}

func BuildMaxHeap(arr []int){
	n := len(arr) - 1
	//the elements in the subarray arr[n/2 + 1 : n] are all leaves of the tree and so each
	//is a 1-element heap to begin with. We go through the remaining of the array from
	//last to first
	
	for i := n/2; i >= 0; i-- {
		MaxHeapify(arr, i)
	}
}