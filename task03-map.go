package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int

	for key := range input {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, elem := range keys {
		result = append(result, input[elem])
	}

	return
}
