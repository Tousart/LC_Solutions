"""
:type s: str
:rtype: bool
"""
        

class Solution(object):
    def isPalindrome(self, s):
        res = ''.join([x.lower() for x in s if x.isalnum()])
        return res == res[::-1]
    

sl = Solution()
print(sl.isPalindrome("A man, a plan, a canal: Panama"))