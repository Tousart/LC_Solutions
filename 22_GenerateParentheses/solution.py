"""
:type n: int
:rtype: List[str]
"""


class Solution(object):
    def generateParenthesis(self, n):
        result = []

        self.make_combination(result, n, 0, 0, "")

        return result

    def make_combination(self, result, n, open_parentheses, close_parentheses, combination):
        if open_parentheses > n or close_parentheses > n or open_parentheses < close_parentheses:
            return

        if open_parentheses == n and close_parentheses == n:
            result.append(combination)
            return

        self.make_combination(result, n, open_parentheses + 1, close_parentheses, combination + "(")
        self.make_combination(result, n, open_parentheses, close_parentheses + 1, combination + ")")


sl = Solution()
print(sl.generateParenthesis(3))