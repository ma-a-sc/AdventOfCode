import itertools
from typing import Tuple

# test value is before the :
# operators are always evaluated left to right - no precedence rules
# two different multiple or addition
# a operator position is between the numbers
# the answer is each one that has at least one solution.
# so I only have to find one
test_input = [
    "190: 10 19",
    "3267: 81 40 27",
    "83: 17 5",
    "156: 15 6",
    "7290: 6 8 6 15",
    "161011: 16 10 13",
    "192: 17 8 14",
    "21037: 9 7 18 13",
    "292: 11 6 16 20",
]


def solve_row(s: str) -> Tuple[bool, int]:
    # can it be soveld or not
    solution, rest = s.split(":")
    solution_num = int(solution)
    numbers = [int(x.strip()) for x in rest.split(" ") if x != ""]
    operator_positions = len(numbers) - 1

    possible_operator_combinations = list(
        itertools.product(["+", "*"], repeat=operator_positions)
    )

    for operators in possible_operator_combinations:
        result = 0
        # I think this misses an operand at the end
        for index, operator in enumerate(operators):
            match operator:
                case "+":
                    if index == 0:
                        result = numbers[index] + numbers[index + 1]
                        continue
                    result = result + numbers[index + 1]
                case "*":
                    if index == 0:
                        result = numbers[index] * numbers[index + 1]
                        continue
                    result = result * numbers[index + 1]
        if result == solution_num:
            return True, solution_num

    return False, 0


with open("/Users/markscharmann/AdventOfCode/assets/day_7_2024.txt", "r") as file:
    all_sums = 0
    for row in file:
        solution_found, number = solve_row(row.strip())
        if solution_found:
            all_sums += number

    print(all_sums)

all_sums = 0

for row in test_input:
    print(row)
    solution_found, number = solve_row(row)
    if solution_found:
        all_sums += number

print(all_sums)
