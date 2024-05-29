package binarysearch

func BinarySearch(arr []int, target int) bool {
	found := false
	hi := len(arr)
	lo := 0

	for lo < hi {
		m := lo + (hi-lo)/2
		v := arr[m]
		if v == target {
			found = true
			break
		} else if v > target {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return found
}
