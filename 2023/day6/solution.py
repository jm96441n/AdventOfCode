import sys
from functools import reduce

import ipdb  # type: ignore # noqa


def part1(lines: list[str]):
    times: list[int] = [int(c) for c in lines[0].split()[1:]]
    distances: list[int] = [int(c) for c in lines[1].split()[1:]]
    ways_to_win: list[int] = []

    for i in range(len(times)):
        wins = 0
        dist = distances[i]
        for t in range(1, times[i]):
            cur_dist = t * (times[i] - t)
            if cur_dist > dist:
                wins += 1
        ways_to_win.append(wins)
    print(ways_to_win)
    return reduce((lambda x, y: x * y), ways_to_win)


def part2(lines: list[str]):
    time: int = int("".join(lines[0].split()[1:]))
    distance: int = int("".join(lines[1].split()[1:]))

    wins = 0
    for t in range(1, time):
        cur_dist = t * (time - t)
        if cur_dist > distance:
            wins += 1
    return wins


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
