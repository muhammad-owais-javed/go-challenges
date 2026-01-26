package main

func FilterBySum(arr [][]int, limit int) [][]int {
	arr1 := make([][]int, 0)

	for i := 0 ; i < len(arr); i++ {
		sum := 0
		for j:=0 ; j < i ; j++ {
			sum += arr[i][j]
		}
		if sum >= limit {
			arr1 = append(arr1, arr[i])

		}
	}
	return arr1
}
