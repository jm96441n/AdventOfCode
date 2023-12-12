import sys
from collections import deque

import ipdb  # type: ignore # noqa


def get_neighbors(max_rows, max_cols, row, col):
    row_delta = [-1, 0, 1, 0]
    col_delta = [0, -1, 0, 1]

    for i in range(len(row_delta)):
        new_row = row_delta[i] + row
        new_col = col_delta[i] + col
        if 0 <= new_row < max_rows and 0 <= new_col < max_cols:
            yield (new_row, new_col)


def part1(lines: list[str]):
    lines = [[c for c in line] for line in lines]
    rows_without_galaxies = []
    cols_without_galaxies = []

    for row in range(len(lines)):
        if all(c == "." for c in lines[row]):
            rows_without_galaxies.append(row)

    for col in range(len(lines[0])):
        collection = []
        for row in range(len(lines)):
            collection.append(lines[row][col])
        if all(c == "." for c in collection):
            cols_without_galaxies.append(col)

    galaxies = []
    i = 1
    for row in range(len(lines)):
        for col in range(len(lines[0])):
            if lines[row][col] == "#":
                galaxies.append((row, col))
                lines[row][col] = i
                i += 1
    for row in reversed(rows_without_galaxies):
        for i in range(len(galaxies)):
            if galaxies[i][0] >= row:
                galaxies[i] = (galaxies[i][0] + 1, galaxies[i][1])

    for col in reversed(cols_without_galaxies):
        for i in range(len(galaxies)):
            if galaxies[i][1] >= col:
                galaxies[i] = (galaxies[i][0], galaxies[i][1] + 1)
    i = 0
    total = 0
    while i < len(galaxies) - 1:
        j = i + 1
        while j < len(galaxies):
            total += abs(galaxies[i][0] - galaxies[j][0]) + abs(galaxies[i][1] - galaxies[j][1])
            j += 1
        i += 1

    return total
    """
    lines = [[c for c in line] for line in lines]
    rows_without_galaxies = []
    cols_without_galaxies = []

    for row in range(len(lines)):
        if all(c == "." for c in lines[row]):
            rows_without_galaxies.append(row)

    for col in range(len(lines[0])):
        collection = []
        for row in range(len(lines)):
            collection.append(lines[row][col])
        if all(c == "." for c in collection):
            cols_without_galaxies.append(col)

    for col in reversed(cols_without_galaxies):
        for row in range(len(lines)):
            lines[row].insert(col, ".")

    for row in reversed(rows_without_galaxies):
        new_row = ["." for _ in range(len(lines[0]))]
        lines.insert(row, new_row)
    galaxies = set([])
    i = 1
    for row in range(len(lines)):
        for col in range(len(lines[0])):
            if lines[row][col] == "#":
                galaxies.add((row, col))
                lines[row][col] = i
                i += 1
    i = 1
    distances = {}
    while i < len(galaxies):
        j = i + 1
        while j <= len(galaxies):
            distances[(i, j)] = float("inf")
            j += 1
        i += 1

    for galaxy in galaxies:
        cur_num = lines[galaxy[0]][galaxy[1]]
        seen = set([galaxy])
        q = deque([galaxy])
        cur_dist = 0
        while q:
            n = len(q)
            cur_dist += 1
            for _ in range(n):
                cur_node = q.popleft()
                for nbr in get_neighbors(len(lines), len(lines[0]), cur_node[0], cur_node[1]):
                    if nbr in seen:
                        continue
                    if nbr in galaxies:
                        found_num = lines[nbr[0]][nbr[1]]
                        loc = (cur_num, found_num)
                        if cur_num > found_num:
                            loc = (found_num, cur_num)
                        distances[loc] = min(cur_dist, distances[loc])

                    q.append(nbr)
                    seen.add(nbr)
    return sum(v for _, v in distances.items())
    """


def part2(lines: list[str]):
    lines = [[c for c in line] for line in lines]
    rows_without_galaxies = []
    cols_without_galaxies = []

    for row in range(len(lines)):
        if all(c == "." for c in lines[row]):
            rows_without_galaxies.append(row)

    for col in range(len(lines[0])):
        collection = []
        for row in range(len(lines)):
            collection.append(lines[row][col])
        if all(c == "." for c in collection):
            cols_without_galaxies.append(col)

    galaxies = []
    i = 1
    for row in range(len(lines)):
        for col in range(len(lines[0])):
            if lines[row][col] == "#":
                galaxies.append((row, col))
                lines[row][col] = i
                i += 1

    for row in reversed(rows_without_galaxies):
        for i in range(len(galaxies)):
            if galaxies[i][0] >= row:
                galaxies[i] = (galaxies[i][0] + 999_999, galaxies[i][1])

    for col in reversed(cols_without_galaxies):
        for i in range(len(galaxies)):
            if galaxies[i][1] >= col:
                galaxies[i] = (galaxies[i][0], galaxies[i][1] + 999_999)
    i = 0
    total = 0
    while i < len(galaxies) - 1:
        j = i + 1
        while j < len(galaxies):
            total += abs(galaxies[i][0] - galaxies[j][0]) + abs(galaxies[i][1] - galaxies[j][1])
            j += 1
        i += 1

    return total


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
