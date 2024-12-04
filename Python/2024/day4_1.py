from collections import namedtuple
from typing import NamedTuple

class XMASfinder(object):

    def __init__(self, matrix: list[list[str]]):
        self.matrix: list[list[str]] = matrix
        self.numbers_of_xmas: int = 0
        self._point = namedtuple('Point', 'x y')
        self.cur_pos: NamedTuple = self._point(0, 0)
        self.pos_to_check: list[list[tuple]] = []
        self._prep_position_arrs()

    def _prep_position_arrs(self) -> None:
        new_positions = []
        # left - works
        arr = [self._point(self.cur_pos.x + x, self.cur_pos.y) for x in range(-3, 1)]
        arr.reverse()
        new_positions.append(arr)
        # right
        new_positions.append([self._point(self.cur_pos.x + x, self.cur_pos.y) for x in range(4)])
        # top
        arr = [self._point(self.cur_pos.x, self.cur_pos.y + x) for x in range(-3, 1)]
        arr.reverse()
        new_positions.append(arr)
        # down
        new_positions.append([self._point(self.cur_pos.x, self.cur_pos.y + x) for x in range(4)])
        # top left
        arr = [self._point(self.cur_pos.x + x, self.cur_pos.y + x) for x in range(-3, 1)]
        arr.reverse()
        new_positions.append(arr)
        # top right
        new_positions.append([self._point(self.cur_pos.x - x, self.cur_pos.y + x) for x in range(4)])
        # bottom left
        new_positions.append([self._point(self.cur_pos.x + x, self.cur_pos.y - x) for x in range(4)])
        # bottom right
        new_positions.append([self._point(self.cur_pos.x + x, self.cur_pos.y + x) for x in range(4)])

        self.pos_to_check = new_positions

    def check_around_for_xmas(self):
        for four_pos in self.pos_to_check:
            try:
                possible_xmas = [self.matrix[pos.x][pos.y] for pos in four_pos if pos.x >= 0 and pos.y >= 0]
            except IndexError:
                continue
            if "".join(possible_xmas) == "XMAS":
                self.numbers_of_xmas += 1

    def go_through_martix(self) -> None:

        for index_x, row in enumerate(self.matrix):
            for index_y, entry in enumerate(row):
                if self.matrix[index_x][index_y] == "X":
                    self.cur_pos = self._point(index_x, index_y)
                    self._prep_position_arrs()
                    self.check_around_for_xmas()
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
