package main

func FilterBySum(arr [][]int, limit int) [][]int {

	arr1 := make([][]int, 0)

	for _, currentElements := range arr {
		sum := 0
		for _, newElements := range currentElements {
			sum += newElements
		}
		if sum >= limit {
			arr1 = append(arr1, currentElements)

		}
	}
	return arr1
}

