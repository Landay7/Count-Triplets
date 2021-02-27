package main

import "fmt"

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	count := int64(0)
	count_map := make(map[int64]int64)
	processed_map := make(map[int64]int64)
	for i := 0; i < len(arr); i++ {
		count_map[arr[i]]++
	}
	for _, val := range arr {
		count_map[val]--
		if _, exists := processed_map[val/r]; exists && val%r == 0 {
			if _, exists = count_map[val*r]; exists {
				count += processed_map[val/r] * count_map[val*r]
			}
		}
		processed_map[val]++
	}
	return count
}

func main() {
	r := int64(3)
	arr := []int64{1, 3, 9, 9, 27, 81}

	ans := countTriplets(arr, r)

	fmt.Println(ans)
}
