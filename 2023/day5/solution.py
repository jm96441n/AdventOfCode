import sys
from collections import defaultdict

DIGITS = set([f"{n}" for n in range(0, 10)])


def parse_field(line_num, lines):
    ranges = []
    while line_num < len(lines) and len(lines[line_num]) > 0 and lines[line_num][0].isnumeric():
        destination, source, range_num = [int(c) for c in lines[line_num].split(" ")]
        ranges.append({"source": (source, source + range_num), "dest": (destination, destination + range_num)})
        line_num += 1
    return line_num + 2, ranges


def part1(lines: list[str]):
    wanted_seeds = [int(c) for c in lines[0].split(" ")[1:]]

    line_num = 3
    line_num, seed_to_soil_ranges = parse_field(line_num, lines)
    line_num, soil_to_fertlizer_ranges = parse_field(line_num, lines)
    line_num, fertilizer_to_water_ranges = parse_field(line_num, lines)
    line_num, water_to_light_ranges = parse_field(line_num, lines)
    line_num, light_to_temperature_ranges = parse_field(line_num, lines)
    line_num, temperature_to_humidity_ranges = parse_field(line_num, lines)
    line_num, humidity_to_location_ranges = parse_field(line_num, lines)

    maps = [
        seed_to_soil_ranges,
        soil_to_fertlizer_ranges,
        fertilizer_to_water_ranges,
        water_to_light_ranges,
        light_to_temperature_ranges,
        temperature_to_humidity_ranges,
        humidity_to_location_ranges,
    ]
    min_location = float("inf")
    for seed in wanted_seeds:
        cur_lookup = seed
        for mapping in maps:
            for ranges in mapping:
                if ranges["source"][0] <= cur_lookup <= ranges["source"][1]:
                    cur_lookup = ranges["dest"][0] + (cur_lookup - ranges["source"][0])
                    break
        min_location = min(min_location, cur_lookup)
    return min_location


def part2(lines: list[str]):
    input = [int(c) for c in lines[0].split(" ")[1:]]
    seeds = []

    line_num = 3
    line_num, seed_to_soil_ranges = parse_field(line_num, lines)
    line_num, soil_to_fertlizer_ranges = parse_field(line_num, lines)
    line_num, fertilizer_to_water_ranges = parse_field(line_num, lines)
    line_num, water_to_light_ranges = parse_field(line_num, lines)
    line_num, light_to_temperature_ranges = parse_field(line_num, lines)
    line_num, temperature_to_humidity_ranges = parse_field(line_num, lines)
    line_num, humidity_to_location_ranges = parse_field(line_num, lines)

    maps = [
        seed_to_soil_ranges,
        soil_to_fertlizer_ranges,
        fertilizer_to_water_ranges,
        water_to_light_ranges,
        light_to_temperature_ranges,
        temperature_to_humidity_ranges,
        humidity_to_location_ranges,
    ]
    min_location = float("inf")

    i = 0
    while i < len(input):
        num, range_num = input[i], input[i + 1]
        i += 2
        seeds.append({"start": num, "end": num + range_num})

    min_location = 0
    while min_location < 10_000_000_000:
        idx = len(maps) - 1
        cur_lookup = min_location
        while idx >= 0:
            cur_mapping = maps[idx]
            for mapping in cur_mapping:
                if mapping["dest"][0] <= cur_lookup < mapping["dest"][1]:
                    cur_lookup = mapping["source"][0] + (cur_lookup - mapping["dest"][0])
                    break
            idx -= 1
        for seed_range in seeds:
            if seed_range["start"] <= cur_lookup < seed_range["end"]:
                return min_location
        min_location += 1


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
