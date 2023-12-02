import string

import ipdb

DIGITS = ("1", "2", "3", "4", "5", "6", "7", "8", "9")
WORDS = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

LETTERS = set(string.ascii_lowercase)


def part1(content):
    sum = 0
    first_digit, second_digit = -1, -1
    for char in content:
        if char == "\n":
            if first_digit > 0:
                second_digit = first_digit if second_digit < 0 else second_digit
                sum += (first_digit * 10) + second_digit
            first_digit, second_digit = -1, -1
        elif char in DIGITS:
            if first_digit == -1:
                first_digit = int(char)
            second_digit = int(char)

    return sum

    """
    for line in lines:
        first_digit = -1
        second_digit = -1
        for c in line:
            if c in DIGITS:
                second_digit = int(c)

                if first_digit == -1:
                    first_digit = int(c)
        if second_digit == -1:
            second_digit = first_digit

        sum += (first_digit * 10) + second_digit
        return sum
    """


def part2(content):
    max_len = 5
    min_len = 3
    sum = 0
    first_digit, second_digit = -1, -1
    word = ""
    i = 0
    while i < len(content):
        while i < len(content) and content[i] in LETTERS:
            word += content[i]
            i += 1
        char = content[i]
        j = 0
        while j + min_len <= len(word):
            for diff in range(min_len, max_len + 1):
                if j + diff > len(word):
                    continue
                if word[j : j + diff] in WORDS:
                    if first_digit == -1:
                        first_digit = WORDS[word[j : j + diff]]
                    second_digit = WORDS[word[j : j + diff]]
            j += 1

        if char == "\n":
            if first_digit > 0:
                sum += (first_digit * 10) + second_digit
            first_digit, second_digit = -1, -1
            word = ""
        elif char in DIGITS:
            if first_digit == -1:
                first_digit = int(char)
            second_digit = int(char)
            word = ""
        i += 1

    return sum


def main():
    with open("./input.txt") as f:
        lines = f.read()
        print(part1(lines))
        print(part2(lines))


if __name__ == "__main__":
    main()
