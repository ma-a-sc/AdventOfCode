from collections import defaultdict
#
# first up build the arrays
left_arr = []
right_arr = []
with open("/Users/markscharmann/AdventOfCode/assets/day_1_2024.txt", "r") as file:
    for line in file:
        left, right = (i for i in line.strip().split(" ") if i != "")
        left_arr.append(left)
        right_arr.append(right)

# build up a counter dict of the right array and then use each entry of the left array as a key to quickly get how often it appreas in the right one
d = defaultdict(int)
for entry in right_arr:
    d[entry] += 1

total_similarity = 0
for entry in left_arr:
    total_similarity += (int(entry) * d[entry])
    print(total_similarity)

print(total_similarity)
