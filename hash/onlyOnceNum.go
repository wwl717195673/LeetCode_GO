package hash

func SingleNumber(nums []int) int {
	storeMap := make(map[int]int, 8)
	for _,j := range nums {
		storeMap[j] += 1
		print(j)
		print(":")
		println(storeMap[j])
	}
	for k,v := range storeMap{
		if v == 1 {
			return storeMap[k]
		}
	}
	return 0
}
