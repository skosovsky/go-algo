package main

func main() {
	twoSum([]int{2, 7, 15}, 9) //nolint:gomnd
}

// result: 64ms / 3.6Mb (Beats 99.98%)
func twoSum(nums []int, target int) []int {
	for i := range nums {
		for y := range nums {
			if nums[i]+nums[y] == target {
				if i != y {
					return []int{i, y}
				}
			}
		}
	}
	return []int{}
}
