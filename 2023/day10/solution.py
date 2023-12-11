import sys
from collections import defaultdict

import ipdb  # type: ignore # noqa

NORTH = "up"
SOUTH = "down"
EAST = "east"
WEST = "west"

DIR_MAPPING = {
    NORTH: (-1, 0),
    SOUTH: (1, 0),
    EAST: (0, 1),
    WEST: (0, -1),
}

DIR_SYM_MAPPING = {
    (NORTH, "|"): NORTH,
    (NORTH, "7"): WEST,
    (NORTH, "F"): EAST,
    (SOUTH, "|"): SOUTH,
    (SOUTH, "J"): WEST,
    (SOUTH, "L"): EAST,
    (EAST, "-"): EAST,
    (EAST, "J"): NORTH,
    (EAST, "7"): SOUTH,
    (WEST, "-"): WEST,
    (WEST, "F"): SOUTH,
    (WEST, "L"): NORTH,
}


def traverse_pipe(lines: list[list[str]]) -> set[tuple[int, int]]:
    start_pos = None
    for i in range(len(lines)):
        for j in range(len(lines[0])):
            if lines[i][j] == "S":
                start_pos = (i, j)
                break
        if start_pos != None:
            break

    seen = set([start_pos])
    cur_dir = None
    # check east
    if lines[start_pos[0]][start_pos[1] + 1] in set(["-", "7", "J"]):
        mapping = {
            "-": EAST,
            "7": SOUTH,
            "J": NORTH,
        }
        cur_pos = (start_pos[0], start_pos[1] + 1)
        cur_dir = mapping[lines[cur_pos[0]][cur_pos[1]]]
    # check west
    elif lines[start_pos[0]][start_pos[1] - 1] in set(["-", "F", "L"]):
        mapping = {
            "-": EAST,
            "F": SOUTH,
            "L": NORTH,
        }
        cur_pos = (start_pos[0], start_pos[1] - 1)
        cur_dir = mapping[lines[cur_pos[0]][cur_pos[1]]]
    # check north
    elif lines[start_pos[0] - 1][start_pos[1]] in set(["|", "7", "F"]):
        mapping = {
            "|": NORTH,
            "7": WEST,
            "F": EAST,
        }
        cur_pos = (start_pos[0] - 1, start_pos[1])
        cur_dir = mapping[lines[cur_pos[0]][cur_pos[1]]]
    # check south
    elif lines[start_pos[0] + 1][start_pos] in set(["|", "L", "J"]):
        mapping = {
            "|": NORTH,
            "J": WEST,
            "L": EAST,
        }
        cur_pos = (start_pos[0], start_pos[1] - 1)
        cur_dir = mapping[lines[cur_pos[0]][cur_pos[1]]]
    print(cur_dir)
    while cur_pos != start_pos:
        seen.add(cur_pos)
        pos_modifier = DIR_MAPPING[cur_dir]
        cur_pos = (cur_pos[0] + pos_modifier[0], cur_pos[1] + pos_modifier[1])
        if cur_pos == start_pos:
            break
        cur_sym = lines[cur_pos[0]][cur_pos[1]]
        cur_dir = DIR_SYM_MAPPING[(cur_dir, cur_sym)]
    return seen


def part1(lines: list[str]):
    lines = [[c for c in line] for line in lines]
    seen = traverse_pipe(lines)
    return (len(seen) + 1) // 2


def part2(lines: list[str]):
    lines = [[c for c in line] for line in lines]
    seen = traverse_pipe(lines)
    inner = 0
    inside = False
    mappings = defaultdict(list)
    for pos in seen:
        mappings[pos[0]].append(pos)

    for k, v in mappings.items():
        positions = sorted(v, key=lambda x: x[1])
        intersections = 0
        inside = False
        for idx, pos in enumerate(positions):
            sym = lines[pos[0]][pos[1]]
            if sym in set(["|", "S"]):
                intersections += 2
            elif sym == "-":
                intersections += 0
            elif sym in set(["J", "F"]):
                intersections -= 1
            elif sym in set(["7", "L"]):
                intersections += 1

            if intersections == 2 or intersections == -2:
                intersections = 0
                inside = not inside

            if inside:
                inner += positions[idx + 1][1] - pos[1] - 1
                x = pos[0]
                y = pos[1] + 1
                while y < positions[idx + 1][1]:
                    lines[x][y] = "I"
                    y += 1

    return inner


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
