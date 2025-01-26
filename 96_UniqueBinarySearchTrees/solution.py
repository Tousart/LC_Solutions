"""
:type n: int
:rtype: int
"""


class Solution(object):
    def numTrees(self, n):
        combinations = [1] + n * [0]

        # Комбинации для i-ого количества узлов
        for i in range(1, n + 1):
            # считаем сумму комбинаций для каждого элемента от первого до последнего
            # в заданном количестве узлов, например: если у нас три узла, то
            # будем считать кол-во комбинаций при корне 1, при корне 2 и при корне 3
            # и суммировать эти количества
            for j in range(0, i):
                combinations[i] += combinations[j] * combinations[i - j - 1]

        return combinations[n]



sl = Solution()
print(sl.numTrees(4))