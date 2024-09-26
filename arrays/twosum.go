package arrays

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, e := range nums {
		c := target - e
		if idx, ok := m[c]; ok {
			return []int{idx, i}
		}
		m[e] = i
	}
	return nil
}

func main() {
	res := twoSum([]int{1, 2}, 3)
	fmt.Print(res)
}
