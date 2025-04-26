"""
:type strs: List[str]
:rtype: List[List[str]]
"""


class Solution(object):
    def groupAnagrams(self, strs):
        s = dict()
        for i in strs:
            key = ''.join(sorted(j for j in i))
            s[key] = s.get(key, []) + [i]
        return s.values()
    

sl = Solution()
print(sl.groupAnagrams(["cab","tin","pew","duh","may","ill","buy","bar","max","doc"]))