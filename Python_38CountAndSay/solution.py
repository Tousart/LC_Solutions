"""
:type n: int
:rtype: str
"""

class Solution(object):
    def countAndSay(self, n):
        result = "1"

        for _ in range(n - 1):
            count = 1
            j = 0
            next_str = ""

            while j != len(result):
                if (j + 1) != len(result):
                    if result[j] == result[j + 1]:
                        count += 1
                    else:
                        next_str += (str(count) + result[j])
                        count = 1
                else:
                    next_str += (str(count) + result[j])
                j += 1

            result = next_str

        return result


sl = Solution()
print(sl.countAndSay(6))