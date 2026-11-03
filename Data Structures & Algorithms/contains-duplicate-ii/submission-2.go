func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)
	for i:=0;i<len(nums);i++{
		if val,ok := seen[nums[i]]; ok{
			if AbsoluteDiff(val,i) <= k{
				return true
			}else{
				seen[nums[i]]=i
			}
		}else{
			seen[nums[i]]= i
		}
	}
	return false
}
func AbsoluteDiff(a,b int)int{
	if a <b{
		return b-a
	}else{
		return a-b
	}
}