from collections import namedtuple
from enum import Enum
from typing import NamedTuple, Set, Tuple, List
import math
import copy


class EndOfMatrixFound(Exception): ...


class Direction(Enum):
    TOP = 0
    RIGHT = 1
    BOTTOM = 2
    LEFT = 3


def check_arr_in_arr_number(arr_short: list, arr_long: list) -> int:
    contains = 0
    for x in range(len(arr_long) - len(arr_short)):
        if arr_short == arr_long[x : x + len(arr_short)]:
            contains += 1

    return contains


class LabFloor(object):
    def __init__(self):
        self.matrix: list[list[str]] = []
        self.modified_matrix: list[list[str]] = []
        self.point = namedtuple("point", "x, y")
        # walking pattern of guard: up until it hits something then turn 90 degrees
        self.guard_position: NamedTuple | None = None
        self.not_possible_position: NamedTuple | None = None
        self.guard_facing_direction: Direction = Direction.TOP
        self.positions: Set = set()
        self.route: List = []
        self.min_sub_array_length = 4
        self.possible_opstruction_positions: Set = set()

    def set_guard_init_position(self):
        for y_coordinate, row in enumerate(self.matrix):
            if "^" in row:
                x_coordinate = row.index("^")
                self.not_possible_position = self.point(x=x_coordinate, y=y_coordinate)
                self.guard_position = self.point(x=x_coordinate, y=y_coordinate)

    def check_next_step_possible(self, altered_matrix: bool = False) -> Tuple:
        match self.guard_facing_direction:
            case Direction.TOP:
                pos = self.point(x=self.guard_position.x, y=self.guard_position.y - 1)
                try:
                    if not altered_matrix:
                        block = self.matrix[pos.y][pos.x]
                    else:
                        block = self.modified_matrix[pos.y][pos.x]
                except IndexError:
                    raise EndOfMatrixFound()

                if block == "#":
                    return False, None, Direction.RIGHT

                return True, pos, None

            case Direction.RIGHT:
                pos = self.point(x=self.guard_position.x + 1, y=self.guard_position.y)
                try:
                    if not altered_matrix:
                        block = self.matrix[pos.y][pos.x]
                    else:
                        block = self.modified_matrix[pos.y][pos.x]
                except IndexError:
                    raise EndOfMatrixFound()

                if block == "#":
                    return False, None, Direction.BOTTOM

                return True, pos, None

            case Direction.BOTTOM:
                pos = self.point(x=self.guard_position.x, y=self.guard_position.y + 1)
                try:
                    if not altered_matrix:
                        block = self.matrix[pos.y][pos.x]
                    else:
                        block = self.modified_matrix[pos.y][pos.x]
                except IndexError:
                    raise EndOfMatrixFound()

                if block == "#":
                    return False, None, Direction.LEFT

                return True, pos, None
            case Direction.LEFT:
                pos = self.point(x=self.guard_position.x - 1, y=self.guard_position.y)
                try:
                    if not altered_matrix:
                        block = self.matrix[pos.y][pos.x]
                    else:
                        block = self.modified_matrix[pos.y][pos.x]
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
                next_step_possible, next_point, direction = (
                    self.check_next_step_possible()
                )
                if not next_step_possible:
                    self.set_new_guard_direction(direction)
                else:
                    self.positions.add(self.guard_position)
                    self.set_new_guard_pos(next_point)

            except EndOfMatrixFound:
                self.positions.add(self.guard_position)
                break

    def check_loop(self) -> bool:
        if not len(self.route) >= self.min_sub_array_length * 2:
            return False

        max_sub_array_length = math.floor(len(self.route) / 2)
        # everything until here now makes sense
        for possible_length in range(
            self.min_sub_array_length, max_sub_array_length + 1
        ):
            for start_pos_of_sub_arr, _ in enumerate(self.route):
                end_position = start_pos_of_sub_arr + possible_length
                sub_arr = self.route[start_pos_of_sub_arr:end_position]
                # setting it to 4 did the trick instead of 2? probably some overlap in paths when setting some obstacles
                if check_arr_in_arr_number(sub_arr, self.route) >= 3:
                    return True

        return False

    def traverse_altered_matrix(self) -> bool:
        counter = 0
        while True:
            try:
                next_step_possible, next_point, direction = (
                    self.check_next_step_possible(altered_matrix=True)
                )
                if not next_step_possible:
                    self.set_new_guard_direction(direction)
                    counter += 1
                else:
                    self.route.append(self.guard_position)
                    self.set_new_guard_pos(next_point)
                    counter += 1

                # min length of a an array with a reocurring sub array with length 4
                if counter >= 8:
                    loop = self.check_loop()
                    if loop:
                        return True

            except EndOfMatrixFound:
                # can never be a loop when the end of matrix will be hit
                return False

    def copy_over_init_matrix_and_override_modified(self):
        self.modified_matrix = copy.deepcopy(self.matrix)

    def prep_for_next_obstruction_run(self, point):
        self.copy_over_init_matrix_and_override_modified()
        self.route = []
        self.modified_matrix[point.y][point.x] = "#"
        lab_floor.set_guard_init_position()
        self.set_new_guard_direction(Direction.TOP)

    def calc_number_of_possible_obstructions(self):
        for point in self.positions:
            print(point)
            if point == self.not_possible_position:
                continue
            self.prep_for_next_obstruction_run(point)
            loop_found = self.traverse_altered_matrix()
            if loop_found:
                self.possible_opstruction_positions.add(point)


grid = [
    [".", ".", ".", ".", "#", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "#"],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", "#", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", "#", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", "#", ".", ".", "^", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "#", "."],
    ["#", ".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", "#", ".", ".", "."],
]

with open("/Users/markscharmann/AdventOfCode/assets/day_6_2024.txt", "r") as file:
    lab_floor = LabFloor()
    lab_floor.matrix = grid
    # for row in file:
    #    lab_floor.matrix.append([x for x in row.strip()])

    lab_floor.set_guard_init_position()

    lab_floor.traverse_matrix()
    # the easy way to calculate these would be to take the path the guard takes and place and obstacle at every possible position once, except the initial one and then let the algo run once to see if the guard gets stuck in a loop.
    # very inefficient cause we would do that a few thousand times. But it gets the job done.
    lab_floor.calc_number_of_possible_obstructions()

    print(lab_floor.possible_opstruction_positions)

# too high 5918
# too low 5328
