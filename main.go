package main

import (
	"fmt"

	"github.com/blitzblade/algorithms/functions"
	"github.com/blitzblade/algorithms/morestrings"
)

func main(){
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))


	//heap
	a := []int{ 1,2,3,4,5,6,7,8,9}
	b := functions.GetHeapItem(a, 1)

	fmt.Println(b)

	hp := []int{4,1,3,2,16,9,10,14,8,7}
	functions.BuildMaxHeap(hp)

	fmt.Println(hp)
}