func hasDuplicate(nums []int) bool {
    seen := make (map[int]bool)
	for i :=0;i<len(nums);i++{
		if _,ok := seen[nums[i]];!ok{
			seen[nums[i]]=true
		}else{
			return true
		}
		
	}
	return false
}