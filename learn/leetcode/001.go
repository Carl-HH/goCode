package main

func main() {
	numbs := []int{2, 7, 11, 15}
	target := 9
	twoSum(numbs, target)

}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}
	return nil
}
