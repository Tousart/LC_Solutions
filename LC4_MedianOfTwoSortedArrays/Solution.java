package LC4_MedianOfTwoSortedArrays;

import java.util.Arrays;

public class Solution {
    public static double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int n = nums1.length, m = nums2.length;

        // Делаем бинарный поиск по меньшему массиву, поэтому меняем массивы местами, если nums1 > nums2
        if (n > m)
            return findMedianSortedArrays(nums2, nums1);

        int lo = 0, hi = n;
        while (lo <= hi) {
            int mid1 = (lo + hi) / 2;
            int mid2 = (n + m + 1) / 2 - mid1;

            // Границы первого массива
            int l1 = (mid1 == 0) ? Integer.MIN_VALUE : nums1[mid1 - 1];
            int r1 = (mid1 == n) ? Integer.MAX_VALUE : nums1[mid1];

            // Границы второго массива
            int l2 = (mid2 == 0) ? Integer.MIN_VALUE : nums2[mid2 - 1];
            int r2 = (mid2 == m) ? Integer.MAX_VALUE : nums2[mid2];

            // Если границы массивов пересеклись
            if (l1 <= r2 && l2 <= r1) {

                if ((n + m) % 2 == 0)
                    return (Math.max(l1, l2) + Math.min(r1, r2)) / 2.0;

                else
                    return Math.max(l1, l2);
            }

            // Если в первом массиве элементы должны быть меньше
            if (l1 > r2)
                hi = mid1 - 1;

            // Если в первом массиве элементы должны быть больше
            else
                lo = mid1 + 1;
        }
        return 0;
    }
//        double poz = (double) (nums1.length + nums2.length) / 2;
//        int i = 0, j = 0;
//
//        while(poz > 0) {
//            if ((nums1[i] >= nums2[j])) {
//                middle1 = nums2[j];
//                j++;
//            } else {
//                middle1 = nums1[i];
//                i++;
//            }
//
//            poz--;
//        }
//        return middle;

    public static void main(String[] args) {
        int[] nums1 = {1, 12, 15, 26, 38};
        int[] nums2 = {2, 13, 17, 30, 45, 60};
        System.out.println(findMedianSortedArrays(nums1, nums2));
    }
}