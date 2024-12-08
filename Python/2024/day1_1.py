# read in the file and fill two arrays

# sort the two arrays
# then subtract and take the aboslut from both and fill a new array
# sum up that array
left_arr = []
right_arr = []
with open("/Users/markscharmann/AdventOfCode/assets/day_1_2024.txt", "r") as file:
    for line in file:
        left, right = (i for i in line.strip().split(" ") if i != "")
        left_arr.append(left)
        right_arr.append(right)

left_arr = sorted(left_arr)
right_arr = sorted(right_arr)

total_distance = 0
for left, right in zip(left_arr, right_arr):
    total_distance += abs(int(left) - int(right))

print(total_distance)
