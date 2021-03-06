package internal

func SliceSplit[T any](size int, arr []T) [][]T {
	result := make([][]T, 0, len(arr)/size+1)
	groupNum := len(arr) / size

	begin := 0
	for i := 0; i <= groupNum; i++ {
		end := begin + size

		if begin == len(arr) {
			break
		}
		if begin+size > len(arr) {
			result = append(result, arr[begin:])
			break
		}

		result = append(result, arr[begin:end])
		begin = end
	}
	return result
}
