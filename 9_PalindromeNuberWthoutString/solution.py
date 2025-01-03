class Solution(object):
    def isPalindrome(self, x):
        if x < 0:
            return False

        old_x = x
        new_x = old_x % 10
        old_x //= 10

        while old_x > 0:
            new_x = (new_x * 10) + (old_x % 10)
            old_x //= 10

        return True if new_x == x else False



sl = Solution()
print(sl.isPalindrome(121))