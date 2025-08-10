package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := (m+n+1)/2 - mid1

		l1 := math.MinInt64
		if mid1 > 0 {
			l1 = nums1[mid1-1]
		}

		r1 := math.MaxInt64
		if mid1 < m {
			r1 = nums1[mid1]
		}

		l2 := math.MinInt64
		if mid2 > 0 {
			l2 = nums2[mid2-1]
		}

		r2 := math.MaxInt64
		if mid2 < n {
			r2 = nums2[mid2]
		}

		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 == 0 {
				return float64((maximum(l1, l2) + minimum(r1, r2))) / 2.0
			}
			return float64(maximum(l1, l2))
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	return 0
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
