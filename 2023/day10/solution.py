import sys
from collections import deque

import ipdb  # type: ignore # noqa


def part1(lines: list[str]):
    row_delta = [-1, 0, 1, 0]
    col_delta = [0, -1, 0, 1]

    lines = [[c for c in line] for line in lines]
    start_pos = None
    for i in range(len(lines)):
        for j in range(len(lines[0])):
            if lines[i][j] == "S":
                start_pos = (i, j)
                break
        if start_pos != None:
            break

    next_q = deque([])
    avail_by_dir = {
        (-1, 0): set(["|", "7", "F"]),
        (0, -1): set(["-", "J", "7"]),
        (1, 0): set(["|", "L", "J"]),
        (0, 1): set(["-", "7", "L"]),
    }

    for i in range(len(row_delta)):
        new_row = row_delta[i] + start_pos[0]
        new_col = col_delta[i] + start_pos[1]
        if lines[new_row][new_col] in avail_by_dir[(row_delta[i], col_delta[i])]:
            next_q.append((new_row, new_col))
            break

    seen = set([start_pos])

    while next_q:
        cur_pos = next_q.popleft()
        seen.add(cur_pos)
        if lines[cur_pos[0]][cur_pos[1]] == "-":
            left_row, left_col = cur_pos[0], cur_pos[1] - 1
            right_row, right_col = cur_pos[0], cur_pos[1] + 1
            if lines[left_row][left_col] in set(["-", "7", "J"]):
                next_q.append((left_row, left_col))
            elif lines[right_row][right_col] in set(["-", "L", "F"]):
                next_q.append((left_row, left_col))

    return (len(seen) + 1) // 2


def part2(lines: list[str]):
    lines = [[c for c in line] for line in lines]


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
