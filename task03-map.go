package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys = make([]int, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, value := range keys {
		result = append(result, input[value])
	}
	return result
}
