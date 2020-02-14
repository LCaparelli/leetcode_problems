"""
4. Median of Two Sorted Arrays

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
"""


class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        merged_list = []

        while len(nums1) > 0 or len(nums2) > 0:
            if len(nums1) == 0:
                merged_list.append(nums2.pop(0))
            elif len(nums2) == 0:
                merged_list.append(nums1.pop(0))
            elif nums1[0] > nums2[0]:
                merged_list.append(nums2.pop(0))
            else:
                merged_list.append(nums1.pop(0))

        return float(merged_list[int(len(merged_list) / 2)]) if len(merged_list) % 2 == 1 else float(
            (merged_list[int(len(merged_list) / 2) - 1] + merged_list[int(len(merged_list) / 2)]) / 2)


solution = Solution()
print(solution.findMedianSortedArrays([1, 3], [2]))
print(solution.findMedianSortedArrays([1, 2], [3, 4]))
