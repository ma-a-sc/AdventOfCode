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


class LabFloor(object):
    def __init__(self):
        self.matrix: list[list[str]] = []
        self.point = namedtuple("point", "x, y")
        # walking pattern of guard: up until it hits something then turn 90 degrees
        self.guard_position: NamedTuple | None = None
        self.not_possible_position: NamedTuple | None = None
        self.guard_facing_direction: Direction = Direction.TOP
        self.positions: Set = set()
        self.route: List = []
        self.turn_points: Set = set()
        self.possible_opstruction_positions: Set = set()

    def set_guard_init_position(self):
        for y_coordinate, row in enumerate(self.matrix):
            if "^" in row:
                x_coordinate = row.index("^")
                self.not_possible_position = self.point(x=x_coordinate, y=y_coordinate)
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

    def determine_next_step(self, obstruction) -> Tuple:
        turn_times = 0
        while True:
            match self.guard_facing_direction:
                case Direction.TOP:
                    pos = self.point(
                        x=self.guard_position.x, y=self.guard_position.y - 1
                    )
                    try:
                        block = self.matrix[pos.y][pos.x]
                    except IndexError:
                        raise EndOfMatrixFound()

                    if block == "#" or pos == obstruction:
                        self.guard_facing_direction = Direction.RIGHT
                        turn_times += 1
                        continue
                    return pos, turn_times

                case Direction.RIGHT:
                    pos = self.point(
                        x=self.guard_position.x + 1, y=self.guard_position.y
                    )
                    try:
                        block = self.matrix[pos.y][pos.x]
                    except IndexError:
                        raise EndOfMatrixFound()

                    if block == "#" or pos == obstruction:
                        self.guard_facing_direction = Direction.BOTTOM
                        turn_times += 1
                        continue
                    return pos, turn_times

                case Direction.BOTTOM:
                    pos = self.point(
                        x=self.guard_position.x, y=self.guard_position.y + 1
                    )
                    try:
                        block = self.matrix[pos.y][pos.x]
                    except IndexError:
                        raise EndOfMatrixFound()

                    if block == "#" or pos == obstruction:
                        self.guard_facing_direction = Direction.LEFT
                        turn_times += 1
                        continue

                    return pos, turn_times
                case Direction.LEFT:
                    pos = self.point(
                        x=self.guard_position.x - 1, y=self.guard_position.y
                    )
                    try:
                        block = self.matrix[pos.y][pos.x]
                    except IndexError:
                        raise EndOfMatrixFound()

                    if block == "#" or pos == obstruction:
                        self.guard_facing_direction = Direction.TOP
                        turn_times += 1
                        continue
                    return pos, turn_times
                case _:
                    raise ValueError

    def traverse_altered_matrix(self, obstruction_point) -> bool:
        while True:
            try:
                # check until
                next_point, turn_times = self.determine_next_step(obstruction_point)
                print(next_point, turn_times)
                if turn_times == 1:
                    if self.guard_position in self.turn_points:
                        return True
                    self.turn_points.add(self.guard_position)
                self.set_new_guard_pos(next_point)

            except EndOfMatrixFound:
                return False

    def prep_for_next_obstruction_run(self):
        self.route = []
        lab_floor.set_guard_init_position()
        self.set_new_guard_direction(Direction.TOP)

    def calc_number_of_possible_obstructions(self):
        for point in self.positions:
            print(point)
            if point == self.not_possible_position:
                continue
            self.prep_for_next_obstruction_run()
            loop_found = self.traverse_altered_matrix(point)
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
    print(len(lab_floor.possible_opstruction_positions))

# too high 2228
