from collections import defaultdict, namedtuple
import itertools
from typing import Set


point = namedtuple("point", "y x")
distance = namedtuple("distance", "yd xd")


class FieldProcessor(object):
    def __init__(self, field_data: list[str]):
        self.field_grid: list[list[str]] = FieldProcessor.process_field_data(field_data)
        self.antenna_data: dict[str, list[point]] = self.get_antenna_data()
        self.antinotes: Set[point] = set()

    @staticmethod
    def process_field_data(field_data) -> list[list[str]]:
        return [[y for y in x.strip()] for x in field_data]

    def get_antenna_data(self) -> dict[str, list[point]]:
        # collecting the antennas works as intended
        data = defaultdict(list)

        for y_index, row in enumerate(self.field_grid):
            for x_index, entry in enumerate(row):
                if entry != ".":
                    data[entry].append(point(y=y_index, x=x_index))
        return data

    def determine_points_of_ani_nodes(self):
        for antenna_type, antennas in self.antenna_data.items():
            print(antenna_type, antennas)
            if len(antennas) <= 1:
                continue

            # the permutations are correct
            for refernece_antenna, other_antenna in itertools.combinations(antennas, 2):
                print(refernece_antenna, other_antenna)

                absolut_distance = distance(
                    yd=abs(refernece_antenna.y - other_antenna.y),
                    xd=abs(refernece_antenna.x - other_antenna.x),
                )

                counter = 0
                while True:
                    point_for_reference_antenna = point(
                        y=refernece_antenna.y - (absolut_distance.yd * counter)
                        if refernece_antenna.y < other_antenna.y
                        else refernece_antenna.y + (absolut_distance.yd * counter),
                        x=refernece_antenna.x - (absolut_distance.xd * counter)
                        if refernece_antenna.x < other_antenna.x
                        else refernece_antenna.x + (absolut_distance.xd * counter),
                    )
                    counter += 1
                    try:
                        if (
                            not point_for_reference_antenna.y < 0
                            and not point_for_reference_antenna.x < 0
                        ):
                            print(point_for_reference_antenna)
                            entry = self.field_grid[point_for_reference_antenna.y][
                                point_for_reference_antenna.x
                            ]
                            self.antinotes.add(point_for_reference_antenna)
                    except IndexError:
                        break

                counter = 0
                while True:
                    point_for_other_antenna = point(
                        y=other_antenna.y - (absolut_distance.yd * counter)
                        if other_antenna.y < refernece_antenna.y
                        else other_antenna.y + (absolut_distance.yd * counter),
                        x=other_antenna.x - (absolut_distance.xd * counter)
                        if other_antenna.x < refernece_antenna.x
                        else other_antenna.x + (absolut_distance.xd * counter),
                    )
                    counter += 1
                    try:
                        if (
                            not point_for_other_antenna.y < 0
                            and not point_for_other_antenna.x < 0
                        ):
                            print(point_for_other_antenna)
                            entry = self.field_grid[point_for_other_antenna.y][
                                point_for_other_antenna.x
                            ]
                            self.antinotes.add(point_for_other_antenna)
                    except IndexError:
                        break


with open("/Users/markscharmann/AdventOfCode/assets/day_8_2024.txt", "r") as file:
    input = []
    for row in file:
        input.append(row.strip())
    print(input)
    processor = FieldProcessor(input)
    processor.determine_points_of_ani_nodes()
    print(processor.antinotes)
    print(len(processor.antinotes))

test_input = [
    "............",
    "........0...",
    ".....0......",
    ".......0....",
    "....0.......",
    "......A.....",
    "............",
    "............",
    "........A...",
    ".........A..",
    "............",
    "............",
]

processor = FieldProcessor(test_input)
processor.determine_points_of_ani_nodes()
print(processor.antinotes)
print(len(processor.antinotes))
