package bobbygogo

/*
# Quick Sort

1. number is value of array

2. as is type of sort : asc, desc
*/
func QuickSort[N float64 | float32 | int32 | int64 | int | string](numbers []N, as string) []N {
	newArr := make([]N, len(numbers))

	copy(newArr, numbers)

	sort(newArr, 0, len(numbers)-1, as)

	return newArr
}

func sort[N float64 | float32 | int32 | int64 | int | string](numbers []N, start, end int, as string) {
	if (end - start) < 1 {
		return
	}

	pivot := numbers[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if as == "desc" {
			if numbers[i] > pivot {
				temp := numbers[splitIndex]
				numbers[splitIndex] = numbers[i]
				numbers[i] = temp
				splitIndex++
			}
		} else {
			if numbers[i] < pivot {
				temp := numbers[splitIndex]
				numbers[splitIndex] = numbers[i]
				numbers[i] = temp
				splitIndex++
			}
		}
	}

	numbers[end] = numbers[splitIndex]
	numbers[splitIndex] = pivot

	sort(numbers, start, splitIndex-1, as)
	sort(numbers, splitIndex+1, end, as)
}
