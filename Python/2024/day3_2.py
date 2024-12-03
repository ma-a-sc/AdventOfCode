# how to get mul(x,y) effeciently out of the line?

# the max length of a correct output would be mul(xxx,yyy) so length of 12, could also (xxx,yyy)mul

# the minimum char length of a relevant sub arr is here 4 = do() and the max is 12 mul(123,123)

# I CAN CUT OUT THE SUB ARRAYS FROM don't() until including do() and then just use my function again but for that I need to treat the lines as one big string and not process line by line cause then I could end up with a dont at the end of the line and having no
# effect

def process_line(s: str) -> int:
    line_score = 0
    sub_arrs = []
    # check which sub arrays to check
    # all sub arrays have the important special chars and mul in it
    for x in range(len(s) - 11):
        sub = s[x:x+12]
        # this fails if the end of the string is 12xmul(1,2)) eg. cause it wont ever satisfy this condition but the input does not end with that so I will ignore it.
        if "mul" in sub and "(" in sub and ")" in sub and "," in sub and sub[:3] == "mul":
            sub_arrs.append(sub)
    print(sub_arrs)
    # parse out the subs

    for arr in sub_arrs:
        # this does not work
        if arr[:3] != "mul" or arr[3] != "(":
            continue
        print(arr)
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
        print(left, right)
        line_score += int(left) * int(right)

    return line_score

def pre_process(s: str) -> str:

    split_by_dont = s.split("don't")
    all_enabled = []
    for sub_line in split_by_dont:
        # multiple do's behind each other brick my logic
        do_splits = sub_line.split("do()")
        for right_of_do in do_splits[1:]:
            all_enabled.append(right_of_do)

    print(all_enabled)
    return "".join(all_enabled)

total_product_instructions = 0
with open("/Users/markscharmann/AdventOfCode/assets/day_3_2024.txt", "r") as file:
    big_line = ""
    test_line = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)don't()+mul(32,64](mul(11,8)undo()asdd12do()mul(8,5))123f12a"
    for line in file:
        if "mul" not in line:
            continue
        big_line += line
    pre_line = pre_process("do()" + big_line)
    print(pre_line)
    score = process_line(pre_line)
    print(score)
    total_product_instructions += score

    print(total_product_instructions)

# too low 20835565 111762583
