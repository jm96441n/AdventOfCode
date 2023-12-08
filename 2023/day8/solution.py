import sys
from math import gcd

import ipdb  # type: ignore # noqa


def part1(lines: list[str]):
    mapping = {}
    directions = lines[0]

    for line in lines[2:]:
        first, left, right = line[0:3], line[7:10], line[12:15]
        mapping[first] = [left, right]

    moves = 0
    current = "AAA"
    while current != "ZZZ":
        idx = moves % len(directions)
        if directions[idx] == "L":
            current = mapping[current][0]
        else:
            current = mapping[current][1]
        moves += 1
    return moves


def part2(lines: list[str]):
    directions = lines[0]
    mapping = {}
    starts = []

    for line in lines[2:]:
        first, left, right = line[0:3], line[7:10], line[12:15]
        mapping[first] = [left, right]
        if first[2] == "A":
            starts.append(first)

    moves = [0 for _ in range(len(starts))]
    for idx, current in enumerate(starts):
        while current[2] != "Z":
            dir_idx = moves[idx] % len(directions)
            if directions[dir_idx] == "L":
                current = mapping[current][0]
            else:
                current = mapping[current][1]
            moves[idx] += 1

    lcm = 1
    for i in moves:
        lcm = lcm * i // gcd(lcm, i)
    return lcm


def main():
    if len(sys.argv) < 2:
        raise Exception("Must pass an file argument to the invocation")

    filename = sys.argv[1]
    with open(filename) as f:
        lines = f.read().splitlines()
        # print(part1(lines))
        print(part2(lines))


if __name__ == "__main__":
    main()
