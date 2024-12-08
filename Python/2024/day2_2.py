# if check safe return false build out as many sub arrays as there are entries in which one entry is removed and check if then the report is considered safe
# not very efficient but with this approach I do not have to improve much of the way how I was handling things in the first place and dont have to figure out the
# bad level. My approach was never good at finding the specific bad level


def check_safe(l: list) -> bool:
    diff_arr = [l[x] - l[x + 1] for x in range(len(l) - 1)]
    # no two numbers next to each other are the same
    if 0 in diff_arr:
        return False

    # no step between two numbers are greater than 3
    if any(tuple(abs(x) > 3 for x in diff_arr)):
        return False

    # check if one or the other fits, either ascending or descending
    asc = all(tuple(x < 0 for x in diff_arr))
    des = all(tuple(x > 0 for x in diff_arr))

    # ensure that only asc or des is true
    if asc and des or not asc and not des:
        return False

    return True


total_safe_levels = 0
with open("/Users/markscharmann/AdventOfCode/assets/day_2_2024.txt", "r") as file:
    for line in file:
        l = [int(i) for i in line.strip().split(" ")]
        safe = check_safe(l)
        if not safe:
            safe_arr = []
            for x in range(len(l)):
                c = l.copy()
                c.pop(x)
                safe_arr.append(check_safe(c))

            if any(safe_arr):
                total_safe_levels += 1
        else:
            total_safe_levels += 1

print(total_safe_levels)
