from collections import defaultdict


def order_update(update: tuple[int, ...]) -> tuple[int, ...]:
    # in theory I could just make it simple and inefficient
    # just get all possible orderings of the update and then
    # check which satisfies the conditions
    # but at the end there might be the possibilty of ordering incorrectly then while abbiding the rules?

    # no, get each number then put it in an array in which all the numbers after it are placed, order does not matter
    # Then order all the arrays by their length.
    # That means that the number which must have the most rules applied is the first and so on
    # Then just build up a new arr by getting the first element of each arrays
    # The length of the numbers that must come after them according to the rules basically sorts them into the correct position in the final

    all_arrays_build_by_rule: list[list[int]] = []
    for index, entry in enumerate(update):
        rules_for_entry = page_order_rules[entry]
        all_arrays_build_by_rule.append(
            [entry] + [x for x in rules_for_entry if x in update]
        )

    # sort the arrs build by rule by their length
    all_arrays_build_by_rule.sort(key=len)

    # now get from every list the first entry and build out the correct update
    return tuple(x[0] for x in all_arrays_build_by_rule)


def check_update(update: tuple[int, ...]) -> int:
    middle = 0
    rules_applied: list[bool] = []
    for index, entry in enumerate(update):
        rules_for_entry = page_order_rules[entry]
        rules_applied.append(
            all([update.index(e) >= index for e in rules_for_entry if e in update])
        )

    if not all(rules_applied):
        new_update = order_update(update)
        return new_update[len(new_update) // 2]
    return middle


page_order_rules: dict[int, list[int]] = defaultdict(list)
update_list: list[tuple[int, ...]] = []
total_score_of_middles = 0

with open("/Users/markscharmann/AdventOfCode/assets/day_5_2024.txt", "r") as file:
    updates = False
    for line in file:
        stripped = line
        if not updates:
            try:
                left, right = stripped.split("|")
            except ValueError:
                updates = True
                continue
            page_order_rules[int(left)].append(int(right))
        else:
            update_list.append(tuple(int(x) for x in stripped.split(",")))

for update in update_list:
    total_score_of_middles += check_update(update)
print(total_score_of_middles)
