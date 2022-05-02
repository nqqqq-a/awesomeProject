package Sort

//快速排序
func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot := Partition(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low

}

//冒泡排序
func BubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := i; j < len; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}
