import itertools

class Solution(object):
    def countAndSay(self, n):
        result = "1"

        for _ in range(n - 1):
            result = "".join([(str(len(list(group))) + digit) for digit, group in itertools.groupby(result)])

        return result




sl = Solution()
print(sl.countAndSay(6))