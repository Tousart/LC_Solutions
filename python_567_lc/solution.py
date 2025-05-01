from collections import Counter

class Solution:
    def checkInclusion(self, s1, s2):
        cntr = Counter(s1)
        step = len(s1)
        for i in range(len(s2)):
            if s2[i] in cntr:
                cntr[s2[i]] -= 1
            if i >= step and s2[i - step] in cntr:
                cntr[s2[i - step]] += 1
            # print(cntr)
            if all(j == 0 for j in cntr.values()):
                # print(all([j for j in cntr.values()]))
                return True
        return False
    

sl = Solution()
print(sl.checkInclusion("hello", "ooolleoooleh"))