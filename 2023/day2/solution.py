import sys
from dataclasses import dataclass

import ipdb

digits = set([f"{x}" for x in range(0, 10)])


def update_counts(counts: dict[str, int], pull: str):
    i = 0
    count = ""
    while pull[i] in digits:
        count += pull[i]
        i += 1
    j = i
    while j < len(pull) and pull[j] != "," and pull[j] != ";":
        j += 1
    counts[pull[i:j]] = int(count)


VALID_RED_CUBES = 12
VALID_GREEN_CUBES = 13
VALID_BLUE_CUBES = 14


@dataclass
class Round:
    blue: int
    red: int
    green: int


@dataclass
class Game:
    id: int
    rounds: list[Round]

    def is_valid(self):
        for round in self.rounds:
            if round.red > VALID_RED_CUBES or round.green > VALID_GREEN_CUBES or round.blue > VALID_BLUE_CUBES:
                return False
        return True


def parse_game(game_str) -> int:
    in_game_part, in_round_part = True, False
    game_id, game_id_str, round_str = -1, "", ""
    counts = {"blue": 0, "red": 0, "green": 0}
    rounds: list[Round] = []
    for char in game_str:
        if in_round_part:
            if char == ";":
                update_counts(counts, round_str)
                rounds.append(Round(blue=counts["blue"], red=counts["red"], green=counts["green"]))
                counts = {"blue": 0, "red": 0, "green": 0}
                round_str = ""
            elif char == ",":
                update_counts(counts, round_str)
                rounds.append(Round(blue=counts["blue"], red=counts["red"], green=counts["green"]))
                round_str = ""
            elif char != " ":
                round_str += char
        if in_game_part:
            if char in digits:
                game_id_str += char
            elif char == ":":
                game_id = int(game_id_str)
                in_game_part = False
                in_round_part = True
    update_counts(counts, round_str)
    rounds.append(Round(blue=counts["blue"], red=counts["red"], green=counts["green"]))
    return Game(id=game_id, rounds=rounds)


def part1(content: list[str]) -> int:
    sum = 0
    for game_str in content:
        game = parse_game(game_str)
        sum += game.id if game.is_valid() else 0

    return sum


def part2(content: list[str]):
    sum = 0
    for game_str in content:
        game = parse_game(game_str)
        max_round = Round(blue=0, green=0, red=0)
        for round in game.rounds:
            max_round.red = max(max_round.red, round.red)
            max_round.green = max(max_round.green, round.green)
            max_round.blue = max(max_round.blue, round.blue)
        sum += (max_round.red or 1) * (max_round.blue or 1) * (max_round.green or 1)
    return sum


def main():
    if len(sys.argv) < 2:
        raise Exception("Must pass an file argument to the invocation")

    filename = sys.argv[1]
    with open(filename) as f:
        content = f.read().splitlines()
        print(part1(content))
        print(part2(content))


if __name__ == "__main__":
    main()
