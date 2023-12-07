import sys
from collections import defaultdict
from dataclasses import dataclass

import ipdb  # type: ignore # noqa

FIVE_OF_A_KIND = 6
FOUR_OF_A_KIND = 5
FULL_HOUSE = 4
THREE_OF_A_KIND = 3
TWO_PAIR = 2
ONE_PAIR = 1
HIGH_CARD = 0

CARDS_1 = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2", "1"]
CARD_VALUES_1 = {k: len(CARDS_1) - i for i, k in enumerate(CARDS_1)}

CARDS_2 = ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "1", "J"]
CARD_VALUES_2 = {k: len(CARDS_1) - i for i, k in enumerate(CARDS_2)}


@dataclass
class Hand:
    cards: list[str]
    kind: int
    bid: int
    card_values: dict[str, int]

    def __gt__(self, other):
        if self.kind > other.kind:
            return True

        if self.kind < other.kind:
            return False

        for i in range(len(self.cards)):
            if self.card_values[self.cards[i]] > self.card_values[other.cards[i]]:
                return True

            if self.card_values[self.cards[i]] < self.card_values[other.cards[i]]:
                return False

    def __lt__(self, other):
        if self.kind < other.kind:
            return True

        if self.kind > other.kind:
            return False

        for i in range(len(self.cards)):
            if self.card_values[self.cards[i]] < self.card_values[other.cards[i]]:
                return True

            if self.card_values[self.cards[i]] > self.card_values[other.cards[i]]:
                return False


def kind_for_cards_2(cards: list[str]) -> int:
    counts = defaultdict(int)
    joker_count = 0
    for c in cards:
        if c == "J":
            joker_count += 1
            continue
        counts[c] += 1

    nums = defaultdict(int)
    three_count_seen = False
    two_count_seen = False
    two_pair_seen = False
    for _, v in counts.items():
        nums[v] += 1

    if joker_count == 0:
        return kind_for_cards_1(cards)

    if joker_count == 1:
        if 4 in nums:
            return FIVE_OF_A_KIND
        if 3 in nums:
            return FOUR_OF_A_KIND
        if 2 in nums and nums[2] == 2:
            return FULL_HOUSE
        if 2 in nums:
            return THREE_OF_A_KIND
        return ONE_PAIR

    if joker_count == 2:
        if 3 in nums:
            return FIVE_OF_A_KIND
        if 2 in nums:
            return FOUR_OF_A_KIND
        return THREE_OF_A_KIND

    if joker_count == 3:
        if 2 in nums:
            return FIVE_OF_A_KIND
        return FOUR_OF_A_KIND

    if joker_count == 4 or joker_count == 5:
        return FIVE_OF_A_KIND


def kind_for_cards_1(cards: list[str]) -> int:
    counts = defaultdict(int)
    for c in cards:
        counts[c] += 1

    nums = defaultdict(int)
    three_count_seen = False
    two_count_seen = False
    for _, v in counts.items():
        nums[v] += 1

    for k, v in nums.items():
        if k == 5:
            return FIVE_OF_A_KIND
        if k == 4:
            return FOUR_OF_A_KIND
        if k == 3:
            three_count_seen = True
        if k == 2 and v == 2:
            return TWO_PAIR
        if k == 2:
            two_count_seen = True

    if three_count_seen and two_count_seen:
        return FULL_HOUSE

    if three_count_seen:
        return THREE_OF_A_KIND

    if two_count_seen:
        return ONE_PAIR

    return HIGH_CARD


def parse_into_hands(lines: list[str], kind_fn, card_values) -> list[Hand]:
    hands = []
    for l in lines:
        cards, bid = l.split()
        cards = [*cards]
        kind = kind_fn(cards)
        hands.append(Hand(cards=[*cards], kind=kind, bid=int(bid), card_values=card_values))
    return hands


def merge_sort(hands: list[Hand]) -> list[Hand]:
    if len(hands) <= 1:
        return

    mid = len(hands) // 2
    left = hands[0:mid]
    right = hands[mid:]

    merge_sort(left)
    merge_sort(right)

    i = j = k = 0
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            hands[k] = left[i]
            i += 1
        else:
            hands[k] = right[j]
            j += 1
        k += 1

    while i < len(left):
        hands[k] = left[i]
        i += 1
        k += 1

    while j < len(right):
        hands[k] = right[j]
        j += 1
        k += 1


def part1(lines: list[str]):
    hands = parse_into_hands(lines, kind_for_cards_1, CARD_VALUES_1)
    merge_sort(hands)
    total = 0
    for i, hand in enumerate(hands):
        total += hand.bid * (i + 1)
    return total


def part2(lines: list[str]):
    hands = parse_into_hands(lines, kind_for_cards_2, CARD_VALUES_2)
    merge_sort(hands)
    total = 0
    for i, hand in enumerate(hands):
        total += hand.bid * (i + 1)
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
