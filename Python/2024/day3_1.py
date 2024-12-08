# how to get mul(x,y) effeciently out of the line?

# the max length of a correct output would be mul(xxx,yyy) so length of 12, could also (xxx,yyy)mul


def process_line(s: str) -> int:
    line_score = 0
    sub_arrs = []
    # check which sub arrays to check
    # all sub arrays have the important special chars and mul in it
    for x in range(len(s) - 12):
        sub = s[x : x + 12]

        if "mul" in sub and "(" in sub and ")" in sub and "," in sub:
            sub_arrs.append(sub)

    # parse out the subs

    for arr in sub_arrs:
        if arr[:3] != "mul" or arr[3] != "(":
            continue
        try:
            numbers, _ = arr[4:].split(")")
        except ValueError:
            try:
                numbers, _, _ = arr[4:].split(")")
            except ValueError:
                continue

        try:
            left, right = numbers.split(",")
        except ValueError:
            continue
        try:
            int(left)
            int(right)
        except:
            continue

        line_score += int(left) * int(right)

    return line_score


total_product_instructions = 0
with open("/Users/markscharmann/AdventOfCode/assets/day_3_2024.txt", "r") as file:
    for line in file:
        if "mul" not in line:
            continue
        score = process_line(line)
        print(score)
        total_product_instructions += score

    print(total_product_instructions)
