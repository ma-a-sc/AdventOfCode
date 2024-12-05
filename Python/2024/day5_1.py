from collections import defaultdict

def check_update(update: tuple[int, ...]) -> int:
    middle = 0
    rules_applied: list[bool] = []
    for index, entry in enumerate(update):
        rules_for_entry = page_order_rules[entry]
        rules_applied.append(all([update.index(e) >= index for e in rules_for_entry if e in update]))


    if all(rules_applied):
        return update[len(update)//2]
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
