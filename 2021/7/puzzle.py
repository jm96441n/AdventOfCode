def run():
    with open("./input.txt") as f:
        positions = [int(i) for i in f.readline().split(",")]
    print(get_distance(positions, _sum_of_distance_from_point))
    print(get_distance(positions, _weighted_sum_of_distance_from_point))


def _sum_of_distance_from_point(input: list[int], point: int) -> int:
    distance = 0
    for i in input:
        distance += abs(i - point)
    return distance


def _weighted_sum_of_distance_from_point(input: list[int], point: int) -> int:
    distance = 0
    for i in input:
        cur_dist = abs(i - point)
        distance += cur_dist + sum(range(0, cur_dist))
    return distance


def get_distance(input: list[int], dist_calc: callable) -> int:
    pos = sum(input) // len(input)
    cur_distance = dist_calc(input, pos)
    next_inc = dist_calc(input, pos + 1)
    next_dec = dist_calc(input, pos - 1)
    change = 1
    next_dist = next_inc
    if next_dec < cur_distance:
        change = -1
        next_dist = next_dec
    pos += change
    while next_dist < cur_distance:
        cur_distance = next_dist
        pos += change
        next_dist = dist_calc(input, pos)

    return cur_distance


if __name__ == "__main__":
    run()
