package homework

func reverse(input []int64) (result []int64) {
	var lastNdx int = len(input) - 1
	for i := lastNdx; i >= 0; i-- {
		result = append(result, input[i])
	}
	return
}
