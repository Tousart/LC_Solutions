"""
:type num: int
:rtype: str
"""


class Solution(object):
    def __init__(self):
        self.letters = {1: "I",
              5: "V",
              10: "X",
              50: "L",
              100: "C",
              500: "D",
              1000: "M"}

    def intToRoman(self, num):
        result = ""

        k = len(str(num))

        for i in str(num):
            category = 10 ** (k - 1)
            k -= 1
            digit = int(i)

            if digit == 1 or digit == 2 or digit == 3:
                result += self.letters[category] * digit

            elif digit == 4:
                result += (self.letters[category] + self.letters[category * 5])

            elif digit == 5:
                result += self.letters[category * 5]

            elif digit == 6 or digit == 7 or digit == 8:
                result += (self.letters[category * 5] + self.letters[category] * (digit - 5))

            elif digit == 9:
                result += self.letters[category] + self.letters[category * 10]


        return result



sl = Solution()
print(sl.intToRoman(1994))