"""
:type firstList: List[List[int]]
:type secondList: List[List[int]]
:rtype: List[List[int]]
"""


class Solution(object):
    def intervalIntersection(self, firstList, secondList):
        result = []
        i = 0
        j = 0

        while i < len(firstList) and j < len(secondList):
            a = firstList[i]
            b = secondList[j]
            if a[0] <= b[1] and b[0] <= a[1]:
                result.append([max(a[0], b[0]), min(a[1], b[1])])

            if a[1] >= b[1]:
                j += 1
            else:
                i += 1

        return result


sl = Solution()
print(sl.intervalIntersection([[0, 2], [5, 10], [13, 23], [24, 25]], [[1, 5], [8, 12], [15, 24], [25, 26]]))