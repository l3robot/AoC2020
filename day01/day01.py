from itertools import combinations
from functools import reduce

if __name__ == "__main__":
    with open("day01/data.txt") as f:
        numbers = f.read().split("\n")

    # Part1
    for couple in combinations(numbers, 2):
        couple = [int(x) if len(x) > 0 else 0 for x in couple]
        if sum(couple) == 2020:
            print("Part One")
            print("--------")
            print(*couple)
            print(reduce(lambda x, y: x * y, couple))
            print()
            break

    # Part2
    for triple in combinations(numbers, 3):
        triple = [int(x) if len(x) > 0 else 0 for x in triple]
        if sum(triple) == 2020:
            print("Part Two")
            print("--------")
            print(*triple)
            print(reduce(lambda x, y: x * y, triple))
            print()
            break
