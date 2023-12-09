import sys
from math import gcd

import ipdb  # type: ignore # noqa


def get_next(nums: list[int]) -> int:
    if all(n == 0 for n in nums):
        return 0

    i = 0
    diffs = []
    while i < len(nums) - 1:
        diffs.append(nums[i + 1] - nums[i])
        i += 1

    next_diff = get_next(diffs)

    return nums[-1] + next_diff


def part1(lines: list[str]):
    lines = [[int(c) for c in line.split()] for line in lines]
    sum = 0
    for n in lines:
        sum += get_next(n)
    return sum


def get_prev(nums: list[int]) -> int:
    if all(n == 0 for n in nums):
        return 0

    i = 0
    diffs = []
    while i < len(nums) - 1:
        diffs.append(nums[i + 1] - nums[i])
        i += 1

    next_diff = get_prev(diffs)

    return nums[0] - next_diff


def part2(lines: list[str]):
    lines = [[int(c) for c in line.split()] for line in lines]
    sum = 0
    for n in lines:
        sum += get_prev(n)
    return sum


def main():
    if len(sys.argv) < 2:
        raise Exception("Must pass an file argument to the invocation")

    filename = sys.argv[1]
    with open(filename) as f:
        lines = f.read().splitlines()
        print(part1(lines))
        print(part2(lines))


if __name__ == "__main__":
    main()
