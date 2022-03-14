package utils

type SortUtils struct {
}

//冒泡排序
func (SortUtils) BubbleSort(arr []int) {
	size := len(arr)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
	}
}

//选择排序
func (SortUtils) SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

//插入排序
func (SortUtils) InsertSort(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

//快速排序
func (SortUtils) QuickSort(arr []int, l, r int) {
	if l < r {
		pivot := arr[r]
		i := l - 1
		for j := l; j < r; j++ {
			if arr[j] <= pivot {
				i++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		i++
		arr[r], arr[i] = arr[i], arr[r]
		SortUtils{}.QuickSort(arr, l, i-1)
		SortUtils{}.QuickSort(arr, i+1, r)
	}
}

//合并
func Merge(arr []int, l, mid, r int) {
	// 分别复制左右子数组
	n1, n2 := mid-l+1, r-mid
	left, right := make([]int, n1), make([]int, n2)
	copy(left, arr[l:mid+1])
	copy(right, arr[mid+1:r+1])
	i, j := 0, 0
	k := l
	for ; i < n1 && j < n2; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	for ; i < n1; i++ {
		arr[k] = left[i]
		k++
	}
	for ; j < n2; j++ {
		arr[k] = right[j]
		k++
	}
}

//合并排序
func (SortUtils) MergeSort(arr []int, l, r int) {
	if l < r {
		mid := (l + r - 1) / 2
		SortUtils{}.MergeSort(arr, l, mid)
		SortUtils{}.MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)
	}
}

//堆调整
func adjust_heap(arr []int, i, size int) {
	if i <= (size-2)/2 {
		//左右子节点
		l, r := 2*i+1, 2*i+2
		m := i
		if l < size && arr[l] > arr[m] {
			m = l
		}
		if r < size && arr[r] > arr[m] {
			m = r
		}
		if m != i {
			arr[m], arr[i] = arr[i], arr[m]
			adjust_heap(arr, m, size)
		}
	}
}

//建堆
func build_heap(arr []int) {
	size := len(arr)
	//从最后一个子节点开始向前调整
	for i := (size - 2) / 2; i >= 0; i-- {
		adjust_heap(arr, i, size)
	}
}

//堆排序
func (SortUtils)HeapSort(arr []int) {
	size := len(arr)
	build_heap(arr)
	for i := size - 1; i > 0; i-- {
		//顶部arr[0]为当前最大值,调整到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		adjust_heap(arr, 0, i)
	}
}

