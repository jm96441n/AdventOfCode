from __future__ import annotations


class Cell:
    def __init__(self, val: int):
        self.val = val
        self.marked: bool = False

    def __repr__(self):
        s = f"{self.val}"
        if self.marked:
            s = "\033[91m" + s + "\033[0m"
        return s

    def __eq__(self, other: int):
        return self.val == other


class Board:
    def __init__(self):
        self.cells: dict[tuple[int, int], Cell] = {}
        self.won: bool = False

    def __repr__(self):
        s = ""
        for i in range(0, 5):
            for j in range(0, 5):
                s += f"({i}, {j}): {self.cells[(i, j)]}  "
            s += "\n"
        return s

    @classmethod
    def create_from_rows(cls, rows: list[str]) -> Board:
        board = Board()
        for i, r in enumerate(rows):
            for j, n in enumerate(r.split()):
                board.add_cell(
                    (
                        i,
                        j,
                    ),
                    int(n),
                )
        return board

    def add_cell(self, loc: tuple[int, int], val: int):
        self.cells[loc] = Cell(val)

    def check_num_called(self, val: int):
        for k, v in self.cells.items():
            if val == v:
                v.marked = True
                self._check_for_win(k)

    def _check_for_win(self, loc: tuple(int, int)):
        self.won = self._check_for_vert_win(loc[1]) or self._check_for_horz_win(loc[0])

    def _check_for_vert_win(self, col: int) -> bool:
        all_marked = True
        for row_num in range(0, 5):
            if not self.cells[(row_num, col)].marked:
                all_marked = False
                break
        return all_marked

    def _check_for_horz_win(self, row: int) -> bool:
        all_marked = True
        for col_num in range(0, 5):
            if not self.cells[(row, col_num)].marked:
                all_marked = False
                break
        return all_marked

    def calculate_score(self, n: int) -> int:
        unmarked = 0
        for _, c in self.cells.items():
            if not c.marked:
                unmarked += c.val
        return unmarked * n


def partone(nums: list[int], boards: list[Board]) -> int:
    for n in nums:
        for b in boards:
            b.check_num_called(n)
            if b.won:
                return b.calculate_score(n)


def parttwo(nums: list[int], boards: list[Board]) -> int:

    for n in nums:
        for b in boards:
            b.check_num_called(n)
            if b.won and len(boards) == 1:
                return b.calculate_score(n)
        boards = [board for board in boards if not board.won]


def run():
    nums, boards = _parse_input()

    print(partone(nums, boards))
    print(parttwo(nums, boards))


def _parse_input() -> list[list[int], list[Board]]:
    nums: list[int] = []
    boards: list[Board] = []
    with open("./input.txt", "r") as f:
        rows = []
        for i, l in enumerate(f.readlines()):
            stripped_line = l.strip("\n")
            if i == 0:
                nums = [int(n) for n in stripped_line.split(",")]
            elif len(stripped_line) == 0:
                if len(rows) > 0:
                    boards.append(Board.create_from_rows(rows))
                rows = []
            else:
                rows.append(l)
        if len(rows) > 0:
            boards.append(Board.create_from_rows(rows))

    return nums, boards


if __name__ == "__main__":
    run()
