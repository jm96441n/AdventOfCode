from dataclasses import dataclass
import time


@dataclass
class Fish:
    age: int = 8

    def decrement_age(self):
        self.age -= 1
        if self.age < 0:
            self.age = 6


def partone_naive(fishes: list[Fish]) -> int:
    for i in range(0, 80):
        next_round = []
        new_fishes = []
        for f in fishes:
            if f.age == 0:
                new_fishes.append(Fish())
            f.decrement_age()
            next_round.append(f)
        next_round.extend(new_fishes)
        fishes = next_round.copy()
    return len(fishes)


def partone(fishes: list[int]) -> int:
    return _count(fishes, 80)


def parttwo(fishes: list[int]) -> int:
    return _count(fishes, 256)


def _count(fishes: list[int], days: int) -> int:
    fish_pop = [0] * 9
    for f in fishes:
        fish_pop[f] += 1

    for d in range(0, days):
        new_pop = fish_pop[0]
        fish_pop = fish_pop[1:]
        fish_pop.append(0)
        fish_pop[6] += new_pop
        fish_pop[8] = new_pop

    return sum(fish_pop)


def run():
    with open("./input.txt") as f:
        fishes = [Fish(int(l)) for l in f.readline().split(",")]
        int_fishes = [fish.age for fish in fishes]
    start = time.time()
    partone(int_fishes)
    end = time.time()
    print(f"optimized solution took: {end - start}")

    start = time.time()
    partone_naive(fishes)
    end = time.time()
    print(f"naive solution took: {end - start}")
    # print(parttwo(fishes))


if __name__ == "__main__":
    run()
