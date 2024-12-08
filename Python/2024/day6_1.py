from collections import namedtuple
from enum import Enum
from io import TextIOWrapper
from typing import NamedTuple, Set, Tuple

class EndOfMatrixFound(Exception):
    ...

class Direction(Enum):
    TOP = 0
    RIGHT = 1
    BOTTOM = 2
    LEFT = 3

class LabFloor(object):


    def __init__(self):
        self.matrix: list[list[str]] = []
        self.point = namedtuple("point", "x, y")
        # walking pattern of guard: up until it hits something then turn 90 degrees
        self.guard_position: NamedTuple | None = None
        self.guard_facing_direction: Direction = Direction.TOP
        self.positions: Set = set()


    def set_guard_init_position(self):
        for y_coordinate, row in enumerate(self.matrix):
            if "^" in row:
                x_coordinate = row.index("^")
                self.guard_position = self.point(x=x_coordinate, y=y_coordinate)

    def check_next_step_possible(self) -> Tuple:

        match self.guard_facing_direction:
            case Direction.TOP:
                pos = self.point(x=self.guard_position.x, y=self.guard_position.y - 1)
                try:
                    block = self.matrix[pos.y][pos.x]
                except IndexError:
                    raise EndOfMatrixFound()

                if block == "#":
                    return False, None, Direction.RIGHT

                return True, pos, None

            case Direction.RIGHT:
                pos = self.point(x=self.guard_position.x + 1, y=self.guard_position.y)
                try:
                    block = self.matrix[pos.y][pos.x]
                except IndexError:
                    raise EndOfMatrixFound()

                if block == "#":
                    return False, None, Direction.BOTTOM

                return True, pos, None

            case Direction.BOTTOM:
                pos = self.point(x=self.guard_position.x, y=self.guard_position.y + 1)
                try:
                    block = self.matrix[pos.y][pos.x]
                except IndexError:
                    raise EndOfMatrixFound()

                if block == "#":
                    return False, None, Direction.LEFT

                return True, pos, None
            case Direction.LEFT:
                pos = self.point(x=self.guard_position.x - 1, y=self.guard_position.y)
                try:
                    block = self.matrix[pos.y][pos.x]
                except IndexError:
                    raise EndOfMatrixFound()

                if block == "#":
                    return False, None, Direction.TOP

                return True, pos, None
            case _:
                raise ValueError


    def set_new_guard_pos(self, pos: NamedTuple) -> None:
        self.guard_position = pos

    def set_new_guard_direction(self, direction: Direction) -> None:
        self.guard_facing_direction = direction

    def traverse_matrix(self):

        while True:
            try:
                next_step_possible, next_point, direction = self.check_next_step_possible()
                print(next_step_possible, next_point, direction)
                if not next_step_possible:
                    self.set_new_guard_direction(direction)
                else:
                    self.positions.add(self.guard_position)
                    self.set_new_guard_pos(next_point)

            except EndOfMatrixFound:
                self.positions.add(self.guard_position)
                break

grid = [
                    ['.', '.', '.', '.', '#', '.', '.', '.', '.', '.'],
                    ['.', '.', '.', '.', '.', '.', '.', '.', '.', '#'],
                    ['.', '.', '.', '.', '.', '.', '.', '.', '.', '.'],
                    ['.', '.', '#', '.', '.', '.', '.', '.', '.', '.'],
                    ['.', '.', '.', '.', '.', '.', '.', '#', '.', '.'],
                    ['.', '.', '.', '.', '.', '.', '.', '.', '.', '.'],
                    ['.', '#', '.', '.', '^', '.', '.', '.', '.', '.'],
                    ['.', '.', '.', '.', '.', '.', '.', '.', '#', '.'],
                    ['#', '.', '.', '.', '.', '.', '.', '.', '.', '.'],
                    ['.', '.', '.', '.', '.', '.', '#', '.', '.', '.']
                ]

with open("/Users/markscharmann/AdventOfCode/assets/day_6_2024.txt", "r") as file:
    lab_floor = LabFloor()
    lab_floor.matrix = grid
    #for row in file:
    #    lab_floor.matrix.append([x for x in row.strip()])

    lab_floor.set_guard_init_position()

    lab_floor.traverse_matrix()

    print(lab_floor.positions)
    print(len(lab_floor.positions))
