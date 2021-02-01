package handler

//冒泡排序
func BubbleSort(numbers []int) []int {
	l := len(numbers)
	for ; l > 0; l-- {
		for i := 0; i+1 < l; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}
	return numbers
}

//选择排序
func SelectSort(numbers []int) []int {
	l := len(numbers)
	for ; l > 0; l-- {
		max := 0
		for i := 0; i < l; i++ {
			if numbers[i] > numbers[max] {
				max = i
			}
		}
		if numbers[max] > numbers[l-1] {
			numbers[l-1], numbers[max] = numbers[max], numbers[l-1]
		}
	}
	return numbers
}

//插入排序
func InsertSort(numbers []int) []int {
	l := len(numbers)
	if l < 2 {
		return numbers
	}
	for j := 1; j < l; j++ {
		i := j - 1
		x := numbers[j]
		for ; i >= 0; i-- {
			if numbers[i] > numbers[j] {
				numbers[i+1] = numbers[i]
			}
		}
		numbers[i] = x
	}
	return numbers
}

//归并排序
func MergeSort(numbers []int) []int {
	l := len(numbers)
	if l < 2 {
		return numbers
	}
	a := (l + 1) / 2
	na := MergeSort(numbers[:a])
	nb := MergeSort(numbers[a:])
	r := make([]int, l)
	ai, bi := 0, 0
	al, bl := a, l-a
	for a := 0; a < l; a++ {
		if ai < al && (bi == bl || na[ai] < nb[bi]) {
			r[a] = na[ai]
			ai++
		} else if bi < bl {
			r[a] = nb[bi]
			bi++
		}
	}
	return r
}

func QuickSort(numbers []int) []int {
	l := len(numbers)
	if l < 2 {
		return numbers
	}
	i, j := 0, l-1
	current := numbers[i]
	for j > i {
		if numbers[j] >= current {
			j--
		} else {
			numbers[i] = numbers[j]
			i++
			for i < j && numbers[i] <= current {
				i++
			}
			if numbers[i] > current {
				numbers[j] = numbers[i]
				j--
			}
		}

	}
	numbers[i] = current
	QuickSort(numbers[:i])
	QuickSort(numbers[i+1:])
	return numbers
}
