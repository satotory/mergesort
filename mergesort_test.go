package mergesort

import (
	"fmt"
	"testing"
)

func TestSimpleArr(t *testing.T) {
	arr := []int{1, 5, 6, 3, 2, 5, 8, 9, 100}
	fmt.Println(mergeSort(arr))
}

func TestArrWithNegative(t *testing.T) {
	arr := []int{1, 5, 6, 3, 2, 5, 8, 9, -4, 0, -0, 50, -100}
	fmt.Println(mergeSort(arr))
}
