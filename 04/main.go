package main

func main() {

}

// if len(arr) is even, median = arr[ceil(len(arr)/2)]
// if len(arr) is odd, median = (arr[len(arr)/2] + arr[len(arr)/2 + 1])/2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalSize := len(nums1) + len(nums2)
	arrRes := make([]int, totalSize)
	i, j := 0, 0
	for k := 0; k < totalSize; k++ {
		if i >= len(nums1) {
			arrRes[k] = nums2[j]
			j++
			continue
		}
		if j >= len(nums2) {
			arrRes[k] = nums1[i]
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			arrRes[k] = nums1[i]
			i++
			continue
		}
		arrRes[k] = nums2[j]
		j++
	}

	med := totalSize / 2
	if totalSize%2 == 0 {
		return (float64(arrRes[med-1]) + float64(arrRes[med])) / 2.0
	}
	return float64(arrRes[med])
}
