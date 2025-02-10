"""
:type s: str
:rtype: str
"""

class Solution(object):
    def clearDigits(self, s):
        i = 0
        s_len = len(s)

        while i != s_len:
            if s[i].isdigit():
                s = s[:i - 1] + s[i + 1:]
                i -= 1
                s_len = len(s)

            else:
                i += 1

        return s


sl = Solution()
print(sl.clearDigits("sab34"))