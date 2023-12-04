import sys
from collections import deque
from functools import reduce
from typing import Generator, Tuple

import ipdb


def get_neighbors(row: int, col: int, max_rows: int, max_cols: int) -> Generator[Tuple[int, int], None, None]:
    col_deltas = [-1, -1, 0, 1, 1, 1, 0, -1]
    row_deltas = [0, 1, 1, 1, 0, -1, -1, -1]

    for i in range(0, len(row_deltas)):
        new_row = row_deltas[i] + row
        new_col = col_deltas[i] + col

        if 0 <= new_row < max_rows and 0 <= new_col < max_cols:
            yield new_row, new_col


DIGITS = set([f"{x}" for x in range(0, 10)])


def check_part_num(lines: list[str], num_pos: list[tuple[int, int]]) -> bool:
    max_rows = len(lines)
    max_cols = len(lines[0])
    for n_row, n_col in num_pos:
        for nbr_row, nbr_col in get_neighbors(n_row, n_col, max_rows, max_cols):
            if lines[nbr_row][nbr_col] != "." and lines[nbr_row][nbr_col] not in DIGITS:
                return True
    return False


# improve by only performing DFS around special chars
def part1(lines: list[str]) -> int:
    sum = 0
    max_rows = len(lines)
    max_cols = len(lines[0])
    seen = set([])
    row = 0
    while row < max_rows:
        col = 0
        while col < max_cols:
            num = ""
            num_pos: list[tuple[int, int]] = []
            while col < max_cols and lines[row][col] in DIGITS:
                num += lines[row][col]
                num_pos.append((row, col))
                col += 1

            if len(num) > 0:
                is_part_num = check_part_num(lines, num_pos)
                if is_part_num:
                    sum += int(num)

            col += 1
        row += 1

    return sum


def part2(lines: list[str]) -> int:
    ratios: list[int] = []
    max_rows = len(lines)
    max_cols = len(lines[0])
    row = 0
    while row < max_rows:
        col = 0
        while col < max_cols:
            if lines[row][col] != "*":
                col += 1
                continue
            nbr_nums: list[str] = []

            part_nums = []
            seen: set[tuple[int, int]] = set([])
            for nbr_row, nbr_col in get_neighbors(row, col, max_rows, max_cols):
                if lines[nbr_row][nbr_col] in DIGITS and (nbr_row, nbr_col) not in seen:
                    start_col = nbr_col
                    num = ""
                    while nbr_col < max_cols and lines[nbr_row][nbr_col] in DIGITS:
                        num += lines[nbr_row][nbr_col]
                        seen.add((nbr_row, nbr_col))
                        nbr_col += 1
                    nbr_col = start_col - 1
                    while nbr_col >= 0 and lines[nbr_row][nbr_col] in DIGITS:
                        num = f"{lines[nbr_row][nbr_col]}{num}"
                        seen.add((nbr_row, nbr_col))
                        nbr_col -= 1
                    part_nums.append(int(num))
            if len(part_nums) == 2:
                ratios.append(part_nums[0] * part_nums[1])

            col += 1
        row += 1
    return sum(ratios)


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
