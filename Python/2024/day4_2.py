from collections import namedtuple
from typing import NamedTuple

class XMASfinder(object):

    def __init__(self, matrix: list[list[str]]):
        self.matrix: list[list[str]] = matrix
        self.numbers_of_xmas: int = 0

    def go_through_martix(self) -> None:
        for index_x, row in enumerate(self.matrix):
            for index_y, entry in enumerate(row):
                if self.matrix[index_x][index_y] == "A":
                    try:
                        if not index_x - 1 >= 0:
                            continue

                        if not index_y - 1 >= 0:
                            continue

                        top_left_bottom_right = [self.matrix[index_x - 1][index_y - 1], "A", self.matrix[index_x + 1][index_y + 1]]
                        bottom_left_top_right = [self.matrix[index_x + 1][index_y - 1], "A", self.matrix[index_x - 1][index_y + 1]]
                    except IndexError or AssertionError:
                        continue
                    top_left_bottom_right = "".join(top_left_bottom_right)
                    bottom_left_top_right = "".join(bottom_left_top_right)
                    if top_left_bottom_right in ["SAM", "MAS"] and bottom_left_top_right in ["SAM", "MAS"]:
                        print(index_x, index_y)
                        self.numbers_of_xmas += 1


matrix = []
with open("/Users/markscharmann/AdventOfCode/assets/day_4_2024.txt", "r") as file:
    for line in file:
        # why ever there remains \n with only one strip
        matrix.append([x for x in line.strip().strip()])
test_matrix = [
    ['M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'],
    ['M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'],
    ['A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'],
    ['M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'],
    ['X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'],
    ['X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'],
    ['S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'],
    ['S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'],
    ['M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'],
    ['M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X']
]
finder = XMASfinder(matrix)
finder.go_through_martix()
print(finder.numbers_of_xmas)

# too high 1987
