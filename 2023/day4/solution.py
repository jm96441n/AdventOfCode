import sys
from collections import defaultdict
from dataclasses import dataclass


def num_wins_for_card(winning_set: set[int], played_set: set[int]) -> int:
    total_won = 0
    for n in played_set:
        if n in winning_set:
            total_won += 1
    return total_won


def part1(lines: list[str]) -> int:
    points = 0
    for card in lines:
        card_num, game = card.split(": ")
        winning, played = game.split(" | ")
        winning_set = set([int(c) for c in winning.split()])
        played_set = set([int(c) for c in played.split()])
        total_won = num_wins_for_card(winning_set, played_set)
        if total_won > 0:
            points += 2 ** (total_won - 1)
    return points


@dataclass
class Game:
    id: int
    winning_set: set[int]
    played_set: set[int]


def part2(lines: list[str]) -> int:
    game_counts: dict[int, int] = defaultdict(lambda: 1)
    games: list[Game] = []
    for idx, card in enumerate(lines):
        game_num = idx + 1
        game_counts[game_num]
        card_num, game = card.split(": ")
        winning, played = game.split(" | ")
        games.append(
            Game(
                id=game_num,
                winning_set=set([int(c) for c in winning.split()]),
                played_set=set([int(c) for c in played.split()]),
            )
        )
    for game in games:
        wins_for_game = num_wins_for_card(game.winning_set, game.played_set)
        for inc in range(1, wins_for_game + 1):
            game_counts[game.id + inc] += game_counts[game.id]
    return sum([v for _, v in game_counts.items()])


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
