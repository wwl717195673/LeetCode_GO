package sort

import "fmt"

//func main(){
//	var a = []int{3,4,5,1,2,3,5,6,8,5,3}
//	fmt.Println(a)
//	fmt.Println(len(a)-1)
//	quicksort(a,0,len(a)-1)
//	fmt.Println(a)
//}

func quicksort(arr []int,left int,right int) {

	if left >= right {
		return
	}

	pivot := arr[left]

	leftIndex := left
	rightIndex := right

	s := fmt.Sprintf("left:%d,right%d", left, right)
	fmt.Println(s)
	for left<right{
		for pivot <= arr[right] && left < right {
			right--
		}
		arr[left] = arr[right]
		if pivot >= arr[left] && left < right {
			left++
		}
		arr[right] = arr[left]
		arr[left] = pivot
		quicksort(arr,leftIndex,left-1)
		quicksort(arr,right+1,rightIndex)
	}
}