package warmup

import (
	"fmt"
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {

	arr := []int32{1, 2, 3, 4, 5}

	removedAt1 := removeAtIndex(arr, 2)

	arr2 := []int32{1, 2, 3, 4, 5, 6}
	removedAt := removeAtIndex(arr2, 5)
	//removed2 := removeAtIndex(arr, 1)

	fmt.Println(arr, removedAt1)
	fmt.Println(arr2, removedAt)
}
