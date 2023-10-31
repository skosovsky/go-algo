package main

func main() {
	twoSum([]int{2, 7, 15}, 9)
}

// result: 64ms / 3.6Mb (Beats 99.98%)
func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for y, _ := range nums {
			if nums[i]+nums[y] == target {
				if i != y {
					return []int{i, y}
				}
			}
		}
	}
	return []int{}
}
