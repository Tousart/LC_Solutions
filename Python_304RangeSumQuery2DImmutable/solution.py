"""
:type row1: int
:type col1: int
:type row2: int
:type col2: int
:rtype: int
"""



class NumMatrix(object):

    def __init__(self, matrix):
        self.matrix = matrix

        self.rows = len(matrix)
        self.cols = len(matrix[0])

        self.prefix = [[0] * (self.cols + 1) for _ in range(self.rows + 1)]

        for row in range(self.rows):
            for col in range(self.cols):
                self.prefix[row + 1][col + 1] = (self.matrix[row][col] +
                                                 self.prefix[row][col + 1] +
                                                 self.prefix[row + 1][col] -
                                                 self.prefix[row][col])


    def sumRegion(self, row1, col1, row2, col2):
        return (self.prefix[row2 + 1][col2 + 1] -
                self.prefix[row1][col2 + 1] -
                self.prefix[row2 + 1][col1] +
                self.prefix[row1][col1])



nm = NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]])
print(nm.sumRegion(2, 1, 4, 3))

# temp = nm.prefix
# for i in temp:
#     print(i)