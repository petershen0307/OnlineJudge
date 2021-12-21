package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	// change m,n to index
	for index, m1, n1 := len(nums1)-1, m-1, n-1; index >= 0; index-- {
		// put result at the last index in nums1
		if m1 >= 0 && n1 >= 0 {
			if nums1[m1] > nums2[n1] {
				nums1[index] = nums1[m1]
				m1--
			} else {
				nums1[index] = nums2[n1]
				n1--
			}
			continue
		}
		// only m1
		if m1 >= 0 {
			nums1[index] = nums1[m1]
			m1--
			continue
		}
		// only n1
		if n1 >= 0 {
			nums1[index] = nums2[n1]
			n1--
			continue
		}
	}
}
